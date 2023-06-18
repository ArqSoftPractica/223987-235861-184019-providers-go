package Routes

import (
	"223987-235861-184019-providers/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(r *gin.Engine) {
	grp1 := r.Group("/health")
	{
		grp1.GET("", Controllers.GetHealth)
	}
}
