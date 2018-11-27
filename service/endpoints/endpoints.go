package endpoints

import (
	"context"

	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/service/service"
	"github.com/go-kit/kit/endpoint"
)

type HealthRequest struct {
}

type HealthResponse struct {
	Ok bool
}

type LoadByPhoneRequest struct {
	Phone string
}
type LoadByPhoneResponse struct {
	Customer models.Customer
}

type CreateRequest struct {
	Customer models.Customer
}

type CreateResponse struct {
	Customer models.Customer
	Err      error
}

func makeHealthEndpoint(i service.ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(HealthRequest)
		b := i.Health()
		return HealthResponse{Ok: b}, nil
	}
}

type LoadByIDRequest struct {
	Id uint64
}

type LoadByIDResponse struct {
	Customer models.Customer
	Err      error
}

func makeLoadByIDEndpoint(i service.ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByIDRequest)
		C, err := i.LoadByID(req.Id)
		return LoadByIDResponse{Customer: *C}, err
	}
}

func makeLoadByPhoneEndpoint(i service.ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByPhoneRequest)
		C, err := i.LoadByPhone(req.Phone)
		return LoadByPhoneResponse{Customer: *C}, err
	}
}

func makeCreateEndpoint(i service.ICustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		C, err := i.Create(&req.Customer)
		return CreateResponse{Customer: *C, Err: err}, err
	}
}

type Endpoints struct {
	Health      endpoint.Endpoint
	LoadByID    endpoint.Endpoint
	LoadByPhone endpoint.Endpoint
	Create      endpoint.Endpoint
}

func MakeServerEndpoints(s *service.CustomerAPIService) Endpoints {
	return Endpoints{
		Health:      makeHealthEndpoint(s),
		LoadByID:    makeLoadByIDEndpoint(s),
		LoadByPhone: makeLoadByPhoneEndpoint(s),
		Create:      makeCreateEndpoint(s),
	}
}
