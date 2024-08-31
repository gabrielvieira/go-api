package server

import (
	"context"
	"github.com/gabrielvieira/go-api/internal/api"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/db"
	"github.com/gabrielvieira/go-api/internal/logger"
)

type Server struct {
	api *api.API
}

func New() (*Server, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}

	l, err := logger.New(c)
	if err != nil {
		return nil, err
	}

	d := db.New(c)
	err = d.Open()
	if err != nil {
		return nil, err
	}

	a := api.New(c, d, l)

	return &Server{
		api: a,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.api.Start()
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
