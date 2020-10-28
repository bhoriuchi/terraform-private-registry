package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v4"
)

// ListVersions handles a versions request
func ListVersions(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	ptype := chi.URLParam(r, "type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"namespace": %q, "type": %q}`, namespace, ptype)))
}

// GetDownload gets the provider
func GetDownload(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	ptype := chi.URLParam(r, "type")
	pversion := chi.URLParam(r, "version")
	pos := chi.URLParam(r, "os")
	parch := chi.URLParam(r, "arch")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"namespace": %q, "type": %q, "version": %q, "os": %q, "arch": %q}`, namespace, ptype, pversion, pos, parch)))
}
