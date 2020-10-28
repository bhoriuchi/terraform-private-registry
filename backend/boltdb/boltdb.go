package boltdb

import (
	"fmt"
	"time"

	"github.com/bhoriuchi/terraform-private-registry/backend"
	bolt "go.etcd.io/bbolt"
)

const bucket = "registry"

// Backend implements backend interface
type Backend struct {
	b    *bolt.DB
	path string
}

// Options config configuration options
type Options struct {
	Path string
}

// NewBackend creates a new backend
func NewBackend(options Options) (backend *Backend) {
	return &Backend{
		path: options.Path,
	}
}

// Open initializes the backend db
func (c *Backend) Open() (err error) {
	if len(c.path) == 0 {
		c.path = "registry.db"
	}

	bopts := &bolt.Options{Timeout: 5 * time.Second}
	c.b, err = bolt.Open(c.path, 0600, bopts)
	return
}

// Close closes the
func (c *Backend) Close() (err error) {
	err = c.b.Close()
	return
}

// Get gets a value by key
func (c *Backend) Get(key string) (value []byte, err error) {
	if len(key) == 0 {
		err = fmt.Errorf(backend.ErrKeyNotSpecified)
		return
	}

	err = c.b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = b.Get([]byte(key))

		if value == nil || len(value) == 0 {
			return fmt.Errorf(backend.ErrKeyNotFound)
		}
		return nil
	})
	return
}

// Put puts a key in the store
func (c *Backend) Put(entries []*backend.KeyValue) (err error) {
	if len(entries) == 0 {
		err = fmt.Errorf(backend.ErrKeyNotSpecified)
		return
	}

	for _, entry := range entries {
		if len(entry.Key) == 0 {
			err = fmt.Errorf(backend.ErrKeyNotSpecified)
			return
		}
		if entry.Value == nil || len(entry.Value) == 0 {
			err = fmt.Errorf(backend.ErrValueNotSpecified)
			return
		}

		err = c.b.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			return b.Put([]byte(entry.Key), entry.Value)
		})

		if err != nil {
			return
		}
	}

	return
}

// Del deletes a key in the store
func (c *Backend) Del(keys []string) (err error) {
	if len(keys) == 0 {
		err = fmt.Errorf(backend.ErrKeyNotSpecified)
		return
	}

	for _, key := range keys {
		if len(key) == 0 {
			err = fmt.Errorf(backend.ErrKeyNotSpecified)
			return
		}

		err = c.b.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			return b.Delete([]byte(key))
		})

		if err != nil {
			return
		}
	}

	return
}
