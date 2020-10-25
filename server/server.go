package server

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/bhoriuchi/terraform-private-registry/handlers"
	"github.com/go-chi/chi/v4"
	"github.com/go-chi/chi/v4/middleware"
)

// Server provides all of the settings and method to run a server
type Server struct {
	tlsCert string
	tlsKey  string
	addr    string
}

// Options options for the server
type Options struct {
	TLSCert string
	TLSKey  string
	Addr    string
}

// Start starts the server
func (c *Server) Start() {
	var err error
	var absCert string
	var absKey string
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/.well-known/terraform.json", handlers.HandleDiscovery)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		redirectURL := JoinURL(`https://registry.terraform.io`, r.URL.Path)
		http.Redirect(w, r, redirectURL, 301)
	})

	if absCert, err = filepath.Abs(c.tlsCert); err != nil {
		log.Fatal(err)
	}
	if absKey, err = filepath.Abs(c.tlsKey); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", c.addr)
	log.Fatal(http.ListenAndServeTLS(c.addr, absCert, absKey, r))
}

// NewServer creates a new instance of server
func NewServer(o Options) (s *Server) {
	return &Server{
		tlsCert: o.TLSCert,
		tlsKey:  o.TLSKey,
		addr:    o.Addr,
	}
}

// JoinURL joins a url
func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}
