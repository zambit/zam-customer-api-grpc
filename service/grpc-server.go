package service

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"git.zam.io/microservices/customer-api/pb"
)

type CustomerAPIGRPCServer struct {
	login       grpctransport.Handler
	loadByPhone grpctransport.Handler
	loadByID    grpctransport.Handler
	create      grpctransport.Handler
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger) pb.CustomerAPIServiceGRPCServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &CustomerAPIGRPCServer{
		login:       grpctransport.NewServer(endpoints.Login, grpcDecodeLoginRequest, grpcEncodeLoginResponse, options...),
		loadByID:    grpctransport.NewServer(endpoints.LoadByID, grpcDecodeLoadByIDRequest, grpcEncodeLoadByIDResponse, options...),
		loadByPhone: grpctransport.NewServer(endpoints.LoadByPhone, grpcDecodeLoadByPhoneRequest, grpcEncodeLoginResponse, options...),
		create:      grpctransport.NewServer(endpoints.Create, grpcDecodeCreateRequest, grpcEncodeCreateRequest, options...),
	}
}

func (s *CustomerAPIGRPCServer) Create(ctx context.Context, r *pb.NewCustomerRequest) (*pb.LoadCustomerResponse, error) {
	_, res, err := s.create.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoadCustomerResponse), nil
}

func (s *CustomerAPIGRPCServer) LoadByID(ctx context.Context, r *pb.LoadByIDRequest) (*pb.LoadCustomerResponse, error) {
	_, res, err := s.loadByID.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoadCustomerResponse), nil
}

func (s *CustomerAPIGRPCServer) LoadByPhone(ctx context.Context, r *pb.LoadByPhoneRequest) (*pb.LoadCustomerResponse, error) {
	_, res, err := s.loadByPhone.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoadCustomerResponse), nil
}

func (s *CustomerAPIGRPCServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoadCustomerResponse, error) {
	_, res, err := s.login.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoadCustomerResponse), nil
}
