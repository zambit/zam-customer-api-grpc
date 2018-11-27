package http

import (
	"context"
	"encoding/json"
	"net/http"

	"git.zam.io/microservices/customer-api/service/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/health", httptransport.NewServer(endpoints.Health, DecodeHealthRequest, EncodeHealthResponse))
	m.Handle("/loadbyid", httptransport.NewServer(endpoints.LoadByID, DecodeLoadByIDRequest, EncodeLoadByIDResponse))
	m.Handle("/loadbyphone", httptransport.NewServer(endpoints.LoadByPhone, DecodeLoadByPhoneRequest, EncodeLoadByPhoneResponse))
	m.Handle("/create", httptransport.NewServer(endpoints.Create, DecodeCreateRequest, EncodeCreateResponse))
	return m
}

func DecodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.HealthRequest{}, nil
}

func EncodeHealthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeLoadByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoadByIDRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func EncodeLoadByIDResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeLoadByPhoneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoadByPhoneRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func EncodeLoadByPhoneResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeCreateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
