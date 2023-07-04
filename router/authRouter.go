package router

import (
	"restapi/middleware"
	"restapi/services/controller"

	"github.com/gin-gonic/gin"
)

func apiAuth(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/", controller.Login)
		auth.POST("/otp", controller.OTP)
		auth.PATCH("/logout", middleware.Authentication(), controller.Logout)
	}
}
