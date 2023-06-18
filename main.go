package main

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Routes"
	"223987-235861-184019-providers/Service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	var fileToLoad string
	if os.Getenv("NODE_ENV") == "development" {
		fileToLoad = ".env.development"
	} else {
		fileToLoad = ".env.production"
	}
	err = godotenv.Load(fileToLoad)

	dbConfig := Config.BuildDBConfig()
	dbUrl := Config.DbURL(dbConfig)
	Config.DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if err != nil {
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Provider{}, &Models.Company{})

	err = Config.DB.Exec(`
		ALTER TABLE providers
		ADD CONSTRAINT fk_providers_companies
		FOREIGN KEY (company_id)
		REFERENCES companies(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE;
	`).Error

	if err != nil {
		fmt.Println(err.Error())
	}
	r := gin.Default()
	Routes.SetupProvidersRoutes(r)
	Routes.SetupHealthRoutes(r)
	Routes.SetupAwsUpdateRoutes(r)
	router := r
	port := os.Getenv("PORT")

	// app, err := newrelic.NewApplication(
	// 	newrelic.ConfigAppName("asp-providers"),
	// 	newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	// )

	go func() {
		Service.ReceiveCompanyMessages()
	}()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
