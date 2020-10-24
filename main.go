package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// JoinURL joins a url
func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		redirectURL := JoinURL(`https://registry.terraform.io`, r.URL.Path)
		http.Redirect(w, r, redirectURL, 301)
	})
	log.Println("Listening")
	http.ListenAndServe(":3000", r)
}
