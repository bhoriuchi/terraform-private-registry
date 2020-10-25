package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// Discovery is the payload for discovery
type Discovery struct {
	ModulesV1   string `json:"modules.v1"`
	ProvidersV1 string `json:"providers.v1"`
}

// HandleDiscovery handles discovery
func HandleDiscovery(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, &Discovery{
		ModulesV1:   "/v1/modules/",
		ProvidersV1: "/v1/providers/",
	})
}
