package routers

import (
	"net/http"
	"srv/structs"

	"github.com/gin-gonic/gin"
)

func RegisterRouterPing(router *gin.RouterGroup) {
	router.GET("/v1/ping", ping)
}

//	@BasePath	/api/v1
//	@Summary	ping
//	@Schemes
//	@Description	this for check is alive service
//	@Tags			ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.MessageResponse
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, structs.MessageResponse{Message: "pong"})
}
