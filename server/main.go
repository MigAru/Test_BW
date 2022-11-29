package main

import (
	"fmt"

	"srv/routers"
	"srv/services"
	"github.com/gin-gonic/gin"
)

func registerAPIServices(router *gin.RouterGroup) {
	apiRoute := router.Group("/api")
	routers.RegisterRouterPing(apiRoute)
    routers.RegisterRouterUsers(apiRoute)
}

func main() {
	App := services.CreateApp()
	
	if err := App.Init(); err != nil{
		fmt.Println(err)
		return
	}
	
	router := gin.Default()
	router.Use(gin.Recovery())
	registerAPIServices(&router.RouterGroup)
	go App.Start(router)

	err := App.ShutDown()
	if err != nil {
		fmt.Println(err)
	}
} 
