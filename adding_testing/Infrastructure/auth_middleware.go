package infrastructure

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwtService.ValidateToken(tokenStr)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userId", claims["user_id"])
		c.Next()
	}
}


// func RoleAuthMiddleware(jwtService JWTService, role string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
// 			c.Abort()
// 			return
// 		}
// 		tokenStr := strings.Split(authHeader, "Bearer ")[1]
// 		token, err := jwtService.ValidateToken(tokenStr)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		claims := token.Claims.(jwt.MapClaims)
// 		roleClaim := claims["role"]
// 		if roleClaim != role {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to access this resource"})
// 			c.Abort()
// 			return
// 		}
// 		c.Set("userId", claims["user_id"])
// 		c.Set("role", roleClaim)
// 		c.Next()
// 	}
// }
