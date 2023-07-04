package router

import (
	"restapi/middleware"
	"restapi/services/controller"

	"github.com/gin-gonic/gin"
)

func apiTasks(api *gin.RouterGroup) {
	task := api.Group("/tasks")
	task.Use(middleware.Authentication())
	{
		task.GET("/", controller.GetTasks)
		task.POST("/", controller.CreateTask)
		task.GET("/calenders/:search", controller.GetCalenderWithTask)
		task.GET("/:id", controller.GetTaskByIndex)
		task.PUT("/:id", middleware.TaskAuthorization(), controller.UpdateTask)
		task.PATCH("/:id", middleware.TaskAuthorization(), controller.CompleteTask)
		task.DELETE("/:id", middleware.TaskAuthorization(), controller.DeleteTask)
	}
}
