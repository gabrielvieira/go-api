package api

import (
	"fmt"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	router *gin.Engine
	config *config.Config
	db     *db.DB
	logger *zap.Logger
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func New(config *config.Config, db *db.DB, logger *zap.Logger) *API {
	router := gin.Default()
	router.Use(gin.Recovery()) // panic recovery middleware
	a := &API{
		router: router,
		config: config,
		db:     db,
		logger: logger,
	}
	router.POST("/user", a.CreateUser)
	router.GET("/user/:id", a.GetUser)
	return a
}

func (a *API) Start() error {
	return a.router.Run(fmt.Sprintf("0.0.0.0:%s", a.config.APIPort))
}
