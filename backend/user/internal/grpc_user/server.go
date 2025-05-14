package grpc_user

import (
	"context"
	userv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/user"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/interfase_lib"
	"google.golang.org/grpc"
)

type serverAPI struct {
	userv1.UserServer
	user interfase_lib.User
}

func RegisterServerAPI(gRPC *grpc.Server, user interfase_lib.User) {
	userv1.RegisterUserServer(gRPC, &serverAPI{user: user})
}

func (s *serverAPI) MyUser(ctx context.Context, req *userv1.MyUserRequest) (*userv1.MyUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) MyUserUpdate(ctx context.Context, req *userv1.MyUserUpdateRequest) (*userv1.MyUserUpdateResponse, error) {
	panic("implement me")
}

func (s *serverAPI) User(ctx context.Context, req *userv1.UserRequest) (*userv1.UserResponse, error) {
	panic("implement me")
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
