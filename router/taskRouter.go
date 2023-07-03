package router

import (
	"restapi/controllers"
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

func apiTasks(api *gin.RouterGroup) {
	task := api.Group("/tasks")
	task.Use(middleware.Authentication())
	{
		task.GET("/", controllers.GetTasks)
		task.POST("/", controllers.CreateTask)
		task.GET("/:id", controllers.GetTaskByIndex)
		task.PUT("/:id", middleware.TaskAuthorization(), controllers.UpdateTask)
		task.PATCH("/:id", middleware.TaskAuthorization(), controllers.CompleteTask)
		task.DELETE("/:id", middleware.TaskAuthorization(), controllers.DeleteTask)
	}
}
