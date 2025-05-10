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
	) (userId int64, err error)
	Logout(
		ctx context.Context,
		sessionUuid string,
	) error
}

type SessionResponse struct {
	SessionUUID string `json:"session_uuid"`
	Success     bool   `json:"success"`
	Message     string `json:"message,omitempty"`
}
