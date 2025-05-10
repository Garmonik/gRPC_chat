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
		return nil, status.Error(codes.InvalidArgument, "Login requires email and password")
	}
	LoginResponse, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), req.GetIpAddress())
	if err == nil {
		return &authv1.LoginResponse{
			SessionUuid: LoginResponse.SessionUUID,
		}, nil
	}
	switch LoginResponse.Code {
	case 3:
		return nil, status.Error(codes.InvalidArgument, "Data not valid")
	case 5:
		return nil, status.Error(codes.NotFound, "User not found")
	default:
		return nil, status.Error(codes.Internal, "Error")
	}
}

func (s *serverAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	if req.GetEmail() == "" || req.GetPassword() == "" || req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Register requires email, name and password")
	}
	RegisterResponse, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword(), req.GetName())
	if err == nil {
		return &authv1.RegisterResponse{
			UserId: RegisterResponse.UserId,
		}, nil
	}
	switch RegisterResponse.Code {
	case 3:
		return nil, status.Error(codes.InvalidArgument, "Invalid email, password or name")
	case 6:
		return nil, status.Error(codes.AlreadyExists, "User with this data already exists")
	case 13:
		return nil, status.Error(codes.Internal, "Error creating new user with this data")
	default:
		return nil, status.Error(codes.Internal, "Error creating new user")
	}
}

func (s *serverAPI) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if req.GetSessionUuid() == "" {
		return nil, status.Error(codes.InvalidArgument, "login requires session uuid")
	}
	code := s.auth.Logout(ctx, req.GetSessionUuid(), req.GetUserId())
	switch code {
	case 0:
		return &authv1.LogoutResponse{Message: "The session was closed"}, nil
	case 3:
		return nil, status.Error(codes.InvalidArgument, "The data looks suspicious")
	case 5:
		return nil, status.Error(codes.NotFound, "Session is not available for this user")
	case 7:
		return nil, status.Error(codes.PermissionDenied, "Access denied\n")
	default:
		return nil, status.Error(codes.Internal, "Error logout")
	}
}
