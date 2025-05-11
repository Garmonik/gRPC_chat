package interfase_lib

import "context"

const (
	InvalidArgument  = 3
	NotFound         = 5
	AlreadyExists    = 6
	PermissionDenied = 7
	Internal         = 13
	OK               = 0
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		ipAddress string,
	) (LoginResponse SessionResponse, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
		name string,
	) (RegisterResponse RegisterResponse, err error)
	Logout(
		ctx context.Context,
		sessionUuid string,
		userId string,
	) int
	SessionList(
		ctx context.Context,
		userId string,
	) int
}

type SessionResponse struct {
	SessionUUID string `json:"session_uuid"`
	Success     bool   `json:"success"`
	Code        int    `json:"code"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id"`
	Code   int   `json:"code"`
}
