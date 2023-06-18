package Routes

import (
	"223987-235861-184019-providers/Authorization"
	"223987-235861-184019-providers/Controllers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Authorization.VerifyToken())
	r.Use(Authorization.VerifyRole())
	grp1 := r.Group("/providers")
	{
		grp1.GET("", Controllers.GetProviders)
		grp1.POST("", Controllers.CreateProvider)
		grp1.GET("/:id", Controllers.GetProviderByID)
		grp1.PUT("/:id", Controllers.UpdateProvider)
		grp1.DELETE("/:id", Controllers.DeleteProvider)
	}
	return r
}

func handleRequest(c *gin.Context) {
	_, ok := c.Request.Context().Value("user").(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
}
