package api

import (
	"errors"
	"github.com/gabrielvieira/go-api/internal/db/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	Id string `uri:"id" binding:"required"`
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
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
		return
	}

	var user model.User
	result := a.db.First(&user, req.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &Response{
		Message: "user found with success",
		Data: &GetUserResponse{
			User: user,
		},
	})
}
