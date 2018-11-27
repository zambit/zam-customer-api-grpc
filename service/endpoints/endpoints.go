package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"git.zam.io/microservices/customer-api/service/service"
)

func makeHealthEndpoint(i service.CustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(HealthRequest)
		b := i.Health()
		return HealthResponse{Ok: b}, nil
	}
}

func makeLoadByIDEndpoint(i service.CustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByIDRequest)
		C := i.LoadByID(req.Id)
		return LoadByIDResponse{Customer: *C}, nil
	}
}

func makeLoadByPhoneEndpoint(i service.CustomerAPIService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadByPhoneRequest)
		C := i.LoadByPhone(req.Phone)
		return LoadByPhoneResponse{Customer: *C}, nil
	}
}

func makeCreateEndpoint(i service.CustomerAPIService) endpoint.Endpoint {
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

func MakeServerEndpoints(s service.CustomerAPIService) Endpoints {
	return Endpoints{
		Health:      makeHealthEndpoint(s),
		LoadByID:    makeLoadByIDEndpoint(s),
		LoadByPhone: makeLoadByPhoneEndpoint(s),
		Create:      makeCreateEndpoint(s),
	}
}
