package router

import (
	"time"

	"github.com/jensskott/template-go-microservice/controllers"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/jensskott/template-go-microservice/lib/router/middleware"
	"github.com/sirupsen/logrus"
)

func Setup() *gin.Engine {
	router := gin.New()

	// Set middlewares
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	router.Use(middleware.Cors())

	// Setup health path
	router.GET("/health", controllers.Health)

	// Setup router group
	api := router.Group("/v1/template")

	// Setup routes
	api.GET("/name", controllers.MyApp)

	return router
}
