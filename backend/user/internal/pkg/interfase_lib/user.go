package interfase_lib

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database/models"
)

const (
	InvalidArgument  = 3
	NotFound         = 5
	AlreadyExists    = 6
	PermissionDenied = 7
	Internal         = 13
	OK               = 0
)

type User interface {
	GetMyUserInfo(
		ctx context.Context,
		userId int64,
	) (*models.User, int)
	UserUpdate(
		ctx context.Context,
		userId int64,
		email string,
		bio string) (int, string)
	GetUserInfo(
		ctx context.Context,
		username string,
	) (*models.User, int)
	GetUserList(
		ctx context.Context,
		userId int64,
	) ([]*models.User, int)
	FriendAdd(
		ctx context.Context,
		myUserId int64,
		friendId int64,
	) (int, error)
	FriendDelete(
		myUserId int64,
		friendId int64,
	) (int, error)
	FriendList(
		ctx context.Context,
	) ([]*models.User, int)
	BlockAdd(
		ctx context.Context,
		myUserId int64,
		friendId int64,
	) (int, error)
	BlockDelete(
		myUserId int64,
		friendId int64,
	) (int, error)
	BlockList(
		ctx context.Context,
	) ([]*models.User, int)
}
