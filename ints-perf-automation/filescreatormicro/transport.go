package filescreatormicro

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

//MakeHttpHandler Make http handler
func MakeHttpHandler(ctx context.Context, endpoint Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/create/{type}/{param}").Handler(httptransport.NewServer(
		endpoint.FilesCreateEndpoint,
		decodeFilesCreateRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeFilesCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req FilesCreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.)
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrBadRouting
	}
	vparam, ok := vars["param"]
	if !ok {
		return nil, ErrBadRouting
	}
	return FilesCreateRequest{
		RequestType: requestType,
		Param:       vparam,
	}, nil
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode error
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
