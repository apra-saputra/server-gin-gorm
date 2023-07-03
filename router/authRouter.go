package router

import (
	"restapi/controllers"
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

func apiAuth(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/", controllers.Login)
		auth.POST("/otp", controllers.OTP)
		auth.PATCH("/logout", middleware.Authentication(), controllers.Logout)
	}
}
