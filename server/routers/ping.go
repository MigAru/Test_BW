package routers

import (
	"net/http"
	"srv/structs"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1
// PingExample godoc
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.MessageResponse
//	@Failure		404	{object}	structs.MessageResponse
//	@Router			/ping [get]
func RegisterRouterPing(router *gin.RouterGroup) {
	router.GET("/v1/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, structs.MessageResponse{Message: "pong"})
}
