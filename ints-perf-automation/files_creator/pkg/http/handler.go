package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
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
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeStoreFileMetadataHandler creates the handler logic
func makeStoreFileMetadataHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/store-file-metadata", http1.NewServer(endpoints.StoreFileMetadataEndpoint, decodeStoreFileMetadataRequest, encodeStoreFileMetadataResponse, options...))
}

// decodeStoreFileMetadataResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeStoreFileMetadataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.StoreFileMetadataRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeStoreFileMetadataResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeStoreFileMetadataResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
