package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	User     string `json:"user"   binding:"required"`
	Password string `json:"password"  binding:"required,gte=3,lte=130"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var request LoginRequest
		err := c.ShouldBindBodyWithJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
		fmt.Printf("%+v\n", request)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func LoginHanlder(req LoginRequest) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
