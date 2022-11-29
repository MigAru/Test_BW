package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouterPing(router *gin.RouterGroup) {
	router.GET("/v1/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
}