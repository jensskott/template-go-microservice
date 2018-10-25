package handlers

import (
	"net/http"

	"github.com/getndazn/template-go-microservice/collect"
	"github.com/gin-gonic/gin"
)

func MyApp(c *gin.Context) {
	data := collect.Json()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
