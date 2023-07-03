package middleware

import (
	"errors"
	"net/http"
	"os"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("apikey")

		if apikey == "" {
			utils.JSONResponse(c, http.StatusProxyAuthRequired, nil, "Invalid Api Key", errors.New("Invalid API Key"))
			return
		}

		if apikey != os.Getenv("API_KEY") {
			utils.JSONResponse(c, http.StatusProxyAuthRequired, nil, "Invalid Api Key", errors.New("Invalid API Key"))
			return
		}

		c.Next()
	}
}
