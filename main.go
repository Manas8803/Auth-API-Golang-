package main

import (
	"Gin/Basics/configs"
	docs "Gin/Basics/docs"
	"Gin/Basics/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Registration API
//	@version		1.0
//	@description	This is a registration api for an application.

//	@BasePath	/api/

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	//* Passing the router to all user(auth) routes.
	api := router.Group("/api/v1")
	routes.UserRoute(api)

	//* Connecting to DB
	configs.ConnectDB()

	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:9000")
}
