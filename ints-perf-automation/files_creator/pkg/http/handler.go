package http

import (
	"context"
	"encoding/json"
	"net/http"

	endpoint "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateHandler creates the handler logic
func makeCreateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create", http1.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...))
}

// decodeCreateResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
