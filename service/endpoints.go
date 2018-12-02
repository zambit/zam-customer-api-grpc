package service

import (
	"context"

	"git.zam.io/microservices/customer-api/models"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type HealthRequest struct {
}

type HealthResponse struct {
	Ok bool
}

func makeHealthEndpoint(i ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(HealthRequest)
		b := i.Health(ctx)
		return HealthResponse{Ok: b}, nil
	}
}

type LoadByIDRequest struct {
	ID uint64 `json:"id"`
}

type LoadByIDResponse struct {
	Customer models.Customer `json:"customer"`
	Err      error           `json:"err"`
}

func makeLoadByIDEndpoint(i ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByIDRequest)
		C, err := i.LoadByID(ctx, req.ID)
		return LoadByIDResponse{Customer: *C}, err
	}
}

type LoadByPhoneRequest struct {
	Phone string
}
type LoadByPhoneResponse struct {
	Customer models.Customer
}

func makeLoadByPhoneEndpoint(i ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByPhoneRequest)
		C, err := i.LoadByPhone(ctx, req.Phone)
		return LoadByPhoneResponse{Customer: *C}, err
	}
}

type CreateRequest struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	StatusID   uint64 `json:"status_id"`
	ReferrerID uint64 `json:"referrer_id"`
}

type CreateResponse struct {
	Customer models.Customer `json:"customer"`
	Error    error           `json:"error,omitempty"`
}

func makeCreateEndpoint(i ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		c := &models.Customer{}
		err := i.Create(ctx, &req, c)
		return CreateResponse{Customer: *c, Error: err}, err
	}
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Customer models.Customer `json:"customer"`
	Error    error           `json:"error,omitempty"`
}

func makeLoginEndpoint(i ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		c, err := i.Login(ctx, req.Phone, req.Password)
		return CreateResponse{Customer: *c, Error: err}, err
	}
}

type Endpoints struct {
	Health      *endpoint.Endpoint
	LoadByID    *endpoint.Endpoint
	LoadByPhone *endpoint.Endpoint
	Create      *endpoint.Endpoint
	Login       endpoint.Endpoint
}

func MakeServerEndpoints(s *CustomerAPIService, logger log.Logger) Endpoints {
	var endpHealth endpoint.Endpoint
	{
		endpHealth = makeHealthEndpoint(s)
		// endpHealth = LoggingMiddleware(log.With(logger, "method", "Health"))(endpHealth)
	}

	var endpLoadByID endpoint.Endpoint
	{
		endpLoadByID = makeLoadByIDEndpoint(s)
		endpLoadByID = LoggingMiddleware(log.With(logger, "method", "LoadByID"))(endpLoadByID)
	}

	var endpLoadByPhone endpoint.Endpoint
	{
		endpLoadByPhone = makeLoadByPhoneEndpoint(s)
		endpLoadByPhone = LoggingMiddleware(log.With(logger, "method", "LoadByPhone"))(endpLoadByPhone)
	}

	var endpCreate endpoint.Endpoint
	{
		endpCreate = makeCreateEndpoint(s)
		endpCreate = LoggingMiddleware(log.With(logger, "method", "Create"))(endpCreate)
	}

	var endpLogin endpoint.Endpoint
	{
		endpLogin = makeLoginEndpoint(s)
		endpLogin = LoggingMiddleware(log.With(logger, "method", "Login"))(endpLogin)
	}

	return Endpoints{
		Health:      &endpHealth,
		LoadByID:    &endpLoadByID,
		LoadByPhone: &endpLoadByPhone,
		Create:      &endpCreate,
		Login:       endpLogin,
	}
}
