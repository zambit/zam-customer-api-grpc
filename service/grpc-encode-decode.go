package service

import (
	"context"

	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/pb"
	"github.com/golang/protobuf/ptypes"
)

func grpcDecodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoginRequest)
	return LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	}, nil
}

func grpcEncodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(LoginResponse)
	return convert(&res.Customer), res.Error
}

func grpcDecodeLoadByIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoadByIDRequest)
	return LoadByIDRequest{
		ID: req.Id,
	}, nil
}

func grpcEncodeLoadByIDResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(LoadByIDResponse)
	return convert(&res.Customer), res.Err
}

func grpcDecodeLoadByPhoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoadByPhoneRequest)
	return LoadByPhoneRequest{
		Phone: req.Phone,
	}, nil
}

func grpcEncodeLoadByPhoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(LoadByIDResponse)
	return convert(&res.Customer), res.Err
}

func grpcEncodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(CreateRequest)
	return &pb.NewCustomerRequest{
		Phone:      req.Phone,
		Password:   req.Password,
		StatusId:   req.StatusID,
		ReferrerId: req.ReferrerID,
	}, nil
}

func grpcDecodeCreateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.NewCustomerRequest)
	return CreateRequest{
		Phone:      req.Phone,
		Password:   req.Password,
		StatusID:   req.StatusId,
		ReferrerID: req.ReferrerId,
	}, nil
}

func convert(httpCustomer *models.Customer) *pb.LoadCustomerResponse {
	registeredAt, _ := ptypes.TimestampProto(httpCustomer.RegisteredAt)
	createdAt, _ := ptypes.TimestampProto(httpCustomer.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(httpCustomer.UpdatedAt)
	return &pb.LoadCustomerResponse{
		Id:           httpCustomer.ID,
		Phone:        httpCustomer.Phone,
		StatusId:     httpCustomer.StatusID,
		ReferrerId:   httpCustomer.ReferrerID,
		RegisteregAt: registeredAt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
