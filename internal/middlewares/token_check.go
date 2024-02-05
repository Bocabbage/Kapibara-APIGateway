package middlewares

import (
	"fmt"
	cryptoutils "kapibara-apigateway/internal/crypto"
	"kapibara-apigateway/internal/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenValidationMid() gin.HandlerFunc {
	// Two different token-auth way: (1) cookie (2) auth-header
	return func(c *gin.Context) {
		var tokenString string

		tokenString, err := c.Cookie("_kapibara_access_token")
		if err != nil {
			const BearerSchema = "Bearer "
			authHeader := c.GetHeader("Authorization")
			if len(authHeader) <= len(BearerSchema) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth-token required but not any legal token found."})
				return
			} else if authHeader[:len(BearerSchema)] != BearerSchema {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth-token required but not any legal token found."})
				return
			}
			tokenString = authHeader[len(BearerSchema):]
		}

		token, err := cryptoutils.ParseJWT(tokenString)
		if err != nil {
			// Enhancement: Add context
			logger.Error(fmt.Sprintf("[TokenValidationMid] error: %s", err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid jwt-token."})
			return
		}
		// check exipred
		currentTime := time.Now().Unix()
		if token.Exp < currentTime {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "jwt-token expired."})
			return
		}

		c.Next()
	}
}
