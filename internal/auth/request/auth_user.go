package request

import (
	"jsonrpcmicro/internal/auth/model"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AuthUserRequest struct {
	model.AuthUser
}
