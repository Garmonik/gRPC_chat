package grpc_user

import (
	"context"
	userv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/user"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/interfase_lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	userv1.UserServer
	userServices interfase_lib.User
}

func RegisterServerAPI(gRPC *grpc.Server, user interfase_lib.User) {
	userv1.RegisterUserServer(gRPC, &serverAPI{userServices: user})
}

func (s *serverAPI) MyUser(ctx context.Context, req *userv1.MyUserRequest) (*userv1.MyUserResponse, error) {
	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	user, code := s.userServices.GetMyUserInfo(ctx, req.GetUserId())
	switch code {
	case interfase_lib.OK:
		return &userv1.MyUserResponse{Id: int64(user.ID), Name: user.Name, Email: user.Email, Bio: user.Bio}, nil
	case interfase_lib.NotFound:
		return nil, status.Error(codes.NotFound, "user not found")
	default:
		return nil, status.Error(codes.Internal, "internal server error")
	}
}

func (s *serverAPI) MyUserUpdate(ctx context.Context, req *userv1.MyUserUpdateRequest) (*userv1.MyUserUpdateResponse, error) {
	if req.GetEmail() == "" {
		return &userv1.MyUserUpdateResponse{Message: "incorrect email"}, status.Error(codes.InvalidArgument, "invalid email")
	}
	code, message := s.userServices.UserUpdate(ctx, req.GetUserId(), req.GetBio(), req.GetEmail())
	switch code {
	case interfase_lib.OK:
		return &userv1.MyUserUpdateResponse{Message: message}, nil
	case interfase_lib.NotFound:
		return &userv1.MyUserUpdateResponse{Message: message}, status.Error(codes.NotFound, "user not found")
	case interfase_lib.Internal:
		return &userv1.MyUserUpdateResponse{Message: message}, status.Error(codes.Internal, "internal server error")
	case interfase_lib.InvalidArgument:
		return &userv1.MyUserUpdateResponse{Message: message}, status.Error(codes.InvalidArgument, "invalid argument")
	default:
		return nil, status.Error(codes.Internal, "internal server error")
	}
}

func (s *serverAPI) User(ctx context.Context, req *userv1.UserRequest) (*userv1.UserResponse, error) {
	if req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}
	user, code := s.userServices.GetUserInfo(ctx, req.GetUsername())
	switch code {
	case interfase_lib.OK:
		return &userv1.UserResponse{Id: int64(user.ID), Name: user.Name, Bio: user.Bio}, nil
	case interfase_lib.NotFound:
		return nil, status.Error(codes.NotFound, "user not found")
	default:
		return nil, status.Error(codes.Internal, "internal server error")
	}
}

func (s *serverAPI) UserList(ctx context.Context, req *userv1.UserListRequest) (*userv1.UserListResponse, error) {
	panic("implement me")
}

func (s *serverAPI) FriendAdd(ctx context.Context, req *userv1.FriendAddRequest) (*userv1.FriendAddResponse, error) {
	panic("implement me")
}

func (s *serverAPI) FriendDelete(ctx context.Context, req *userv1.FriendDeleteRequest) (*userv1.FriendDeleteResponse, error) {
	panic("implement me")
}

func (s *serverAPI) FriendList(ctx context.Context, req *userv1.FriendListRequest) (*userv1.FriendListResponse, error) {
	panic("implement me")
}

func (s *serverAPI) BlockAdd(ctx context.Context, req *userv1.BlockAddRequest) (*userv1.BlockAddResponse, error) {
	panic("implement me")
}

func (s *serverAPI) BlockDelete(ctx context.Context, req *userv1.BlockDeleteRequest) (*userv1.BlockDeleteResponse, error) {
	panic("implement me")
}

func (s *serverAPI) BlockList(ctx context.Context, req *userv1.BlockListRequest) (*userv1.BlockListResponse, error) {
	panic("implement me")
}
