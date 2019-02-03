package http

import (
	"context"
	"encoding/json"
	endpoint "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeCreateFilesHandler creates the handler logic
func makeCreateFilesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-files", http1.NewServer(endpoints.CreateFilesEndpoint, decodeCreateFilesRequest, encodeCreateFilesResponse, options...))
}

// decodeCreateFilesResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateFilesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateFilesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateFilesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateFilesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
