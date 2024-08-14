package infrastructure

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



func AuthMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenStr := strings.Split(authHeader,"Bearer ")[1]
		token,err := jwtService.ValidateToken(tokenStr)


		if err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userId",claims["user_id"])
		c.Next()
	}
}



