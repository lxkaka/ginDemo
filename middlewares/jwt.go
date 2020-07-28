package middlewares

import (
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		claims, _ := utils.ParseToken(authHeader)
		if claims == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Set("userid", claims.UserID)
		}
	}
}
