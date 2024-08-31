package server

import (
	"context"
	"github.com/gabrielvieira/go-api/internal/api"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/db"
	"github.com/gabrielvieira/go-api/internal/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	api    *api.API
	config *config.Config
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
		api:    a,
		config: c,
	}, nil
}

func (s *Server) Run(ctx context.Context) {
	// init API and workers
	go func() {
		if err := s.api.Start(ctx); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// wait for quit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// stop API and workers
	if err := s.api.Stop(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
