package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, data interface{}, message string, err error) {
	response := gin.H{
		"status": status,
	}

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(status, gin.H{
			"error":   message,
			"message": "ERROR",
			"status":  response["status"],
		})
	} else {
		if data != nil {
			response["data"] = data
		}
		response["message"] = message
		c.JSON(status, response)
	}

}
