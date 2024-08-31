package api

import (
	"go.uber.org/zap"
)

type User struct {
	logger *zap.Logger
}

func NewUser() *User {
	return &User{}
}
