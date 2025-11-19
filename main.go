package main

import (
	"be/app"
	"be/helper"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	// Env
	port := os.Getenv("APP_PORT")
	jwtKey := os.Getenv("JWT_KEY")

	// Database Config
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db := app.OpenConnection(dbUser, dbPass, dbHost, dbPort, dbName)

	// Other
	validate := validator.New()

	// Dependency
	memberDependency := app.NewMemberDependency(db, validate, jwtKey)

	// Setup Router
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization", "Email", "Code"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	routerParent := app.Router{
		MemberDependency: memberDependency,

		JwtKey: jwtKey,
		Engine: engine,
	}
	router := app.NewRouter(routerParent)

	// Run Server
	if port == "" {
		port = ":3001"
	}
	err = router.Engine.Run(port)
	helper.FatalError(err)
}
