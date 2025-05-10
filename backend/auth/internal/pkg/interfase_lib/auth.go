package interfase_lib

import "context"

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (LoginResponse SessionResponse, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (RegisterResponse RegisterResponse, err error)
	Logout(
		ctx context.Context,
		sessionUuid string,
	) error
}

type SessionResponse struct {
	SessionUUID string `json:"session_uuid"`
	Success     bool   `json:"success"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id"`
	Code   int   `json:"code"`
}
