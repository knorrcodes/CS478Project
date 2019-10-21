package common

import (
	"encoding/json"
	"net/http"

	log "github.com/lfkeitel/verbose/v5"
)

const contentTypeJSON = "application/json; charset=utf-8"

// A APIResponse is returned as a JSON struct to the client.
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewAPIResponse creates an APIResponse object with status c, message m, and data d.
func NewAPIResponse(m string, d interface{}) *APIResponse {
	return &APIResponse{
		Message: m,
		Data:    d,
	}
}

// NewEmptyAPIResponse returns an APIResponse with no message or data.
func NewEmptyAPIResponse() *APIResponse {
	return &APIResponse{}
}

// Encode the APIResponse into JSON.
func (a *APIResponse) Encode() []byte {
	b, err := json.Marshal(a)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"package": "common",
		}).Error("Error encoding API response data")
	}
	return b
}

// WriteResponse encodes and writes a response back to the client.
func (a *APIResponse) WriteResponse(w http.ResponseWriter, code int) (int, error) {
	r := a.Encode()
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return 0, nil
	}
	return w.Write(r)
}

// SystemVersion is the current version of the software.
var SystemVersion string
