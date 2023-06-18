package Routes

import (
	"223987-235861-184019-providers/Authorization"
	"223987-235861-184019-providers/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupAwsUpdateRoutes(r *gin.Engine) {
	grp1 := r.Group("/awsUpdate")
	{
		grp1.Use(Authorization.VerifyToken())
		grp1.Use(Authorization.RoleIsMaster())
		grp1.POST("", Controllers.UpdateAwsCreds)
	}
}
