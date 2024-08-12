package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader:= c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Notoken provided"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader,"Bearer ")
		if authHeader == tokenString {
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token provided"})
			c.Abort()
			return
		}
		token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error) {
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v",token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")),nil
		})
		if claims,ok := token.Claims.(jwt.MapClaims);ok && token.Valid {
			c.Set("user_id",claims["user_id"])
			c.Set("role",claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden,gin.H{"error":"You are not authorized to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}