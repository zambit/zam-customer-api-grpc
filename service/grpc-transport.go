package service

import (
	"context"

	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/pb"
	"github.com/golang/protobuf/ptypes"
)

func grpcDecodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoadByIDRequest)
	return &pb.LoadByIDRequest{
		Id: req.Id,
	}, nil
}

func encodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(LoginResponse)
	return convert(&res.Customer), res.Error
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

// func grpcDecodeLoadByIDRequest(_ context.Context, r interface{}) (interface{}, error) {
// 	req := r.(*pb.LoadByIDRequest)
// 	return &pb.LoadByIDRequest{
// 		Id: req.Id,
// 	}, nil
// }

// func grpcEncodeLoadByIDResponse(_ context.Context, r interface{}) (interface{}, error) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	return json.NewEncoder(w).Encode(response)
// }

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
