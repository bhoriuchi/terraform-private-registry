package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gogo/protobuf/jsonpb"
	pb "github.com/golang/protobuf/proto"
)

var marshaler = jsonpb.Marshaler{}

// Remarshal converts from protobuf to interface
func Remarshal(msg pb.Message) (o interface{}, err error) {
	var str = ""

	if str, err = marshaler.MarshalToString(msg); err != nil {
		return
	}

	err = json.Unmarshal([]byte(str), &o)
	return
}

// JSONResponse renders a json response
func JSONResponse(w http.ResponseWriter, r *http.Request, msg pb.Message) {
	rsp, err := Remarshal(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	render.JSON(w, r, rsp)
}
