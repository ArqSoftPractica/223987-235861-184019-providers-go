package Authorization

import (
	"context"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")[7:] // Remove "Bearer " prefix

		if tokenString != "" {
			publicKeyPath := "./public.key"
			publicKeyBytes, err := ioutil.ReadFile(publicKeyPath)
			if err != nil {
				log.Printf("Failed to read public key: %s", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}

			publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
			if err != nil {
				log.Printf("Failed to parse public key: %s", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			})
			if err != nil {
				log.Printf("Failed to parse token: %s", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(c.Request.Context(), "user", claims)

				c.Request = c.Request.WithContext(ctx)
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided. Auth token is required."})
			return
		}

		c.Next()
	}
}
