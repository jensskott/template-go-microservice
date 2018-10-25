package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "OPTIONS"},
		AllowHeaders:    []string{"Origin"},
		MaxAge:          12 * time.Hour,
	})
}
