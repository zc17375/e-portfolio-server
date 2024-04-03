package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloApi struct{}

func (h *HelloApi) Hello(c *gin.Context) {
	helloService.Hello(c)
}
