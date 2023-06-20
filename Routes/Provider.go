package Routes

import (
	"223987-235861-184019-providers/Authorization"
	"223987-235861-184019-providers/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupProvidersRoutes(r *gin.Engine) {
	grp1 := r.Group("/providers")
	{
		grp1.Use(Authorization.VerifyToken())
		grp1.Use(Authorization.VerifyRole())
		grp1.GET("", Controllers.GetProviders)
		grp1.POST("", Controllers.CreateProvider)
		grp1.GET("/:id", Controllers.GetProviderByID)
		grp1.PUT("/:id", Controllers.UpdateProvider)
		grp1.DELETE("/:id", Controllers.DeleteProvider)
	}
}
