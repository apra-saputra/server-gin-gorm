package middleware

import (
	"errors"
	"net/http"
	"restapi/initializers"
	"restapi/services/model"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func TaskAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReq, _ := c.Get("user")

		user := userReq.(model.User)

		var task model.Task
		result := initializers.DB.First(&task, c.Param("id"))
		if result.RowsAffected == 0 {
			utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", errors.New("Task not found"))
			return
		} else if result.Error != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to fetch task", errors.New("Failed to fetch task"))
			return
		}

		if user.Role == model.Users && user.ID != task.UserID {
			utils.JSONResponse(c, http.StatusForbidden, nil, "Invalid Authorization", errors.New("Invalid Authorization"))
			return
		}

		c.Next()
	}
}
