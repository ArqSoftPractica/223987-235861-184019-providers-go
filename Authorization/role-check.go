package Authorization

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Request.Context().Value("user")
		if user == nil {
			errorMessage := "Unauthorized. You do not have the correct permissions for this action."
			log.Println(errorMessage)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMessage})
			return
		}

		claims, ok := user.(jwt.MapClaims)
		if !ok {
			errorMessage := "Unauthorized. You do not have the correct permissions for this action."
			log.Println(errorMessage)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMessage})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			errorMessage := "Unauthorized. You do not have the correct permissions for this action."
			log.Println(errorMessage)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMessage})
			return
		} else {
			if role != "" {
				rolesWithAccess := map[string]string{
					"provider":    "PROVIDERS",
					"maintenance": "MAINTENANCE",
				}

				if role == "MASTER" || role == "ADMIN" || rolesWithAccess[role] != "" {
					c.Next()
					return
				} else {
					errorMessage := "Unauthorized. You do not have the correct permissions for this action."
					log.Println(errorMessage)
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMessage})
					return
				}
			} else {
				errorMessage := "Unauthorized."
				log.Println(errorMessage)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMessage})
				return
			}
		}
	}
}
