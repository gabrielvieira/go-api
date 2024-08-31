package main

import (
	"fmt"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/logger"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	User     string `json:"user"   binding:"required"`
	Password string `json:"password"  binding:"required,gte=3,lte=130"`
}

func main() {

	c, err := config.New()
	if err != nil {
		panic(err)
	}

	l, err := logger.New(c)
	if err != nil {
		panic(err)
	}

	l.Info("ola")
	fmt.Printf("%+v\n", c)

	//r := gin.Default()
	//r.POST("/login", func(c *gin.Context) {
	//	var request LoginRequest
	//	err := c.ShouldBindBodyWithJSON(&request)
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//	c.Status(http.StatusOK)
	//	fmt.Printf("%+v\n", request)
	//})
	//
	//r.Run() // listen and serve on 0.0.0.0:8080
}

func LoginHanlder(req LoginRequest) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
