package backend

// Backend implements a backend for storing persistent data
type Backend interface {
	// Opens a connection to the backend
	Open() (err error)

	// Close closes connection to the backend
	Close() (err error)

	// Get gets a keys value
	Get(key string) (value []byte, err error)

	// Put puts one or more values into the store
	Put(entries []*KeyValue) (err error)

	// Del deletes one or more values in the store
	Del(keys []string) (err error)
}

// KeyValue a key value entry
type KeyValue struct {
	Key   string
	Value []byte
}
