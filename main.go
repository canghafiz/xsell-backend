package be

import (
	"be/app"
	"be/helper"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	helper.FatalError(err)

	if os.Getenv("APP_STATUS") == "Debug" {
		gin.SetMode(gin.DebugMode)
	}

	if os.Getenv("APP_STATUS") == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Database Config
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db := app.OpenConnection(dbUser, dbPass, dbHost, dbPort, dbName)
}
