package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helpers.ExtractToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := helpers.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)

		c.Next()
	}
}
