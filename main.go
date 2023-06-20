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
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func NewRelicErrorLogger(app *newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path)
		defer txn.End()

		c.Next()
		status := c.Writer.Status()
		if status >= 400 {
			defer txn.End()
			txn.NoticeError(fmt.Errorf("HTTP %d: %s", status, http.StatusText(status)))
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("asp-providers"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(nrgin.Middleware(app))
	r.Use(NewRelicErrorLogger(app))
	Routes.SetupProvidersRoutes(r)
	Routes.SetupHealthRoutes(r)
	Routes.SetupAwsUpdateRoutes(r)
	router := r
	port := os.Getenv("PORT")

	go func() {
		Service.ReceiveCompanyMessages()
	}()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
