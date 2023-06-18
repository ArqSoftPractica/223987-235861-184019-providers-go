package Controllers

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
	Available                   bool   `json:"available"`
	CompanyQueueServiceIsActive bool   `json:"companyQueueServiceIsActive"`
	DbConnection                bool   `json:"dbConnection"`
	Version                     string `json:"version"`
}

func GetHealth(c *gin.Context) {
	var health Health

	health.Available = true
	health.CompanyQueueServiceIsActive = Service.CompanyQueueServiceActive.IsActive
	health.DbConnection = IsDBConnectionAvailable()
	health.Version = "1.0.0"

	c.JSON(http.StatusOK, health)
}

func IsDBConnectionAvailable() bool {
	sqlDB, err := Config.DB.DB()
	if err != nil {
		fmt.Printf("Error accessing *sql.DB: %s\n", err.Error())
		return false
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s\n", err.Error())
		return false
	}

	return true
}
