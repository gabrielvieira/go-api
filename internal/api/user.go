package api

import (
	"github.com/gabrielvieira/go-api/internal/db/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct {
	model.User
}

func (a *API) Login(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (a *API) CreateUser(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (a *API) GetUser(c *gin.Context) {
	c.Status(http.StatusOK)
}
