package handlers

import (
	"net/http"

	"github.com/bhoriuchi/terraform-private-registry/helpers"
	proto "github.com/bhoriuchi/terraform-private-registry/proto"
)

// GetDiscovery handles discovery
func GetDiscovery(w http.ResponseWriter, r *http.Request) {
	helpers.JSONResponse(w, r, &proto.Discovery{
		ModulesV1:   "/v1/modules/",
		ProvidersV1: "/v1/providers/",
	})
}
