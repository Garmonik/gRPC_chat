package grpc_auth

import (
	"context"
	authinterfase "github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/interfase_lib"
	authv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	authv1.UnimplementedAuthServer
	auth authinterfase.Auth
}

func RegisterServerAPI(gRPC *grpc.Server, auth authinterfase.Auth) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "login requires email and password")
	}
	LoginResponse, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err == nil {
		return &authv1.LoginResponse{
			SessionUuid: LoginResponse.SessionUUID,
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "Data not valid")
}

func (s *serverAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "register requires email and password")
	}
	RegisterResponse, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword(), req.GetName())
	return &authv1.RegisterResponse{
		UserId: 123,
	}, nil
}

func (s *serverAPI) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if req.GetSessionUuid() == "" {
		return nil, status.Error(codes.InvalidArgument, "login requires session uuid")
	}
	return &authv1.LogoutResponse{
		Message: "Logout",
	}, nil
}
