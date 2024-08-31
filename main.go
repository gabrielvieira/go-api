package main

import (
	"context"
	"github.com/gabrielvieira/go-api/internal/server"
)

type LoginRequest struct {
	User     string `json:"user"   binding:"required"`
	Password string `json:"password"  binding:"required,gte=3,lte=130"`
}

func main() {
	ctx := context.Background()
	s := server.New()
}
