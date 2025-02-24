package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github/com/Gajju8989/Auth_Service/internal/config/jwtkey"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		var tokenString string

		if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
			tokenString = authHeader[len(bearerPrefix):]
		} else {
			tokenString = authHeader
		}

		jwtKey, err := jwtkey.GetJWTKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Key not found"})
		}

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}

func RefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("R-Token")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		var tokenString string

		if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
			tokenString = authHeader[len(bearerPrefix):]
		} else {
			tokenString = authHeader
		}

		claims := &jwt.StandardClaims{}

		jwtKey, err := jwtkey.GetJWTKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Key not found"})
		}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}
