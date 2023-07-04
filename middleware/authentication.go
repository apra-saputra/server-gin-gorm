package middleware

import (
	"net/http"
	"restapi/initializers"
	"restapi/services/model"
	"restapi/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// mendapat cookies authorization
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			utils.JSONResponse(c, http.StatusUnauthorized, nil, "Invalid Authentication", err)
			return
		}

		// decode token
		token, err := utils.DecodeToken(tokenString)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// validate expired Token
			if float64(time.Now().Unix()) > claims["expIn"].(float64) {
				utils.JSONResponse(c, http.StatusUnauthorized, nil, "Invalid Authentication", err)
				return
			}
			// find user
			var user model.User
			initializers.DB.First(&user, claims["contain"])

			if user.ID == 0 {
				utils.JSONResponse(c, http.StatusUnauthorized, nil, "Invalid Authentication", err)
				return
			}

			// set user
			c.Set("user", user)

			c.Next()
		} else {
			utils.JSONResponse(c, http.StatusUnauthorized, nil, "Invalid Authentication", err)
			return
		}

	}
}
