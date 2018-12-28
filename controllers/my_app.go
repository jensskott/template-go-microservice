package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jensskott/template-go-microservice/lib/collect"
)

func MyApp(c *gin.Context) {
	data := collect.Json()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
