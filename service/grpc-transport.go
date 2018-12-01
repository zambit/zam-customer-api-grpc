package service

import (
	"context"

	"git.zam.io/microservices/customer-api/pb"
)

func grpcEncodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(CreateRequest)
	return &pb.NewCustomerRequest{
		Phone:    req.Phone,
		Password: req.Password,
		StatusId: 1,
	}, nil
}

func grpcDecodeCreateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.NewCustomerRequest)
	return CreateRequest{
		Phone: req.Phone,
	}, nil
}
