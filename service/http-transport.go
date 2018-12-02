package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(endpoints Endpoints, logger log.Logger) http.Handler {
	m := chi.NewMux()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorLogger(logger),
	}

	m.Method(http.MethodGet, "/health", httptransport.NewServer(endpoints.Health, httpDecodeHealthRequest, httpEncodeHealthResponse, options...))
	m.Method(http.MethodPost, "/load-by-id", httptransport.NewServer(endpoints.LoadByID, httpDecodeLoadByIDRequest, httpEncodeLoadByIDResponse, options...))
	m.Method(http.MethodPost, "/loadbyphone", httptransport.NewServer(endpoints.LoadByPhone, httpDecodeLoadByPhoneRequest, httpEncodeLoadByPhoneResponse, options...))
	m.Method(http.MethodPost, "/create", httptransport.NewServer(endpoints.Create, httpDecodeCreateRequest, httpEncodeCreateResponse, options...))
	m.Method(http.MethodPost, "/login", httptransport.NewServer(endpoints.Login, httpDecodeLoginRequest, httpEncodeLoginResponse, options...))

	return m
}

func httpDecodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return HealthRequest{}, nil
}

func httpEncodeHealthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func httpDecodeLoadByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req LoadByIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func httpEncodeLoadByIDResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func httpDecodeLoadByPhoneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req LoadByPhoneRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func httpEncodeLoadByPhoneResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func httpDecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func httpEncodeCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func httpDecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func httpEncodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"error"`
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
