package router

import (
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.Use(middleware.ApiKey())
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Server Ready...! 🚀",
			})
		})

		apiAuth(api)
		apiTasks(api)
	}

	return router
}
