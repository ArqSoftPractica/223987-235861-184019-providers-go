package Routes

import (
	"223987-235861-184019-providers/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
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
