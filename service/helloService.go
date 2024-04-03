package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloService struct{}

func (h *HelloService) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}