package service

import (
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
