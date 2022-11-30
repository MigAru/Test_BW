package main

import (
	"fmt"

	docs "srv/docs"
	"srv/middlewares"
	"srv/routers"
	"srv/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerAPIServices(router *gin.RouterGroup) {
	apiRoute := router.Group("/api")
	apiRoute.Use(middlewares.CheckJSONApplicationMiddleware())
	routers.RegisterRouterPing(apiRoute)
	routers.RegisterRouterUsers(apiRoute)
	routers.RegisterRouterTransactions(apiRoute)
}

func main() {
	App := services.CreateApp()

	if err := App.Init(); err != nil {
		fmt.Println(err)
		return
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	registerAPIServices(&router.RouterGroup)
	go App.Start(router)

	err := App.ShutDown()
	if err != nil {
		fmt.Println(err)
	}
}
