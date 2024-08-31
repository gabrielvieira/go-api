package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type API struct {
	httpServer *http.Server
	router     *gin.Engine
	config     *config.Config
	db         *db.DB
	logger     *zap.Logger
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func New(config *config.Config, db *db.DB, logger *zap.Logger) *API {
	router := gin.Default()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.APIPort),
		Handler: router.Handler(),
	}
	a := &API{
		httpServer: httpServer,
		router:     router,
		config:     config,
		db:         db,
		logger:     logger,
	}
	a.SetupRoutes()
	return a
}

func (a *API) SetupRoutes() {
	a.router.POST("/user", a.CreateUser)
	a.router.GET("/user/:id", a.GetUser)
}

func (a *API) Start(ctx context.Context) error {
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (a *API) Stop(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
