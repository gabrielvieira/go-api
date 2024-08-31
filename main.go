package main

import (
	"context"
	"github.com/gabrielvieira/go-api/internal/server"
)

func main() {
	ctx := context.Background()
	s, err := server.New()
	if err != nil {
		panic(err)
	}
	s.Run(ctx)
}
