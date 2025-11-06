package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// ✅ Debug log
		fmt.Println("JWT Claims =>", claims)

		// ✅ Handle all possible claim types
		isAdminBool := false
		if val, exists := claims["isAdmin"]; exists {
			switch v := val.(type) {
			case bool:
				isAdminBool = v
			case string:
				isAdminBool = (v == "true" || v == "1")
			case float64:
				isAdminBool = v == 1
			default:
				isAdminBool = false
			}
		}

		if !isAdminBool {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not an admin!"})
			c.Abort()
			return
		}

		c.Next()
	}
}
