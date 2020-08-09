package router

import (
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(200, "404")
	})
	return r
}
