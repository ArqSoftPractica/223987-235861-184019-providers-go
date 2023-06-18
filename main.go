package main

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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
	Config.DB, err = gorm.Open("mysql", dbUrl)

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Provider{}, &Models.Company{})

	err = Config.DB.Model(&Models.Provider{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		panic(err)
	}

	router := Routes.SetupRouter()
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
