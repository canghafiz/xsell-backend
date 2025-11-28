package main

import (
	"be/apps"
	"be/dependencies"
	"be/helpers"
	"be/models/domains"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	helpers.FatalError(err)

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
	db := apps.OpenConnection(dbUser, dbPass, dbHost, dbPort, dbName)

	// Redis Config
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisConfig := domains.RedisConfig{
		Prefix:   os.Getenv("REDIS_PREFIX"),
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:       redisDb,
		Password: os.Getenv("REDIS_PASS"),
	}

	// Twilio Config
	twilio := domains.Twilio{
		AccountId:  os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
	}

	// Other
	validate := validator.New()

	// Dependency
	memberDependency := dependencies.NewMemberDependency(db, validate, redisConfig, jwtKey, twilio)
	adminDependency := dependencies.NewAdminDependency(db, validate)
	dependency := dependencies.NewDependency()

	// Setup Router
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization", "Email", "Code"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	routerParent := apps.Router{
		MemberDependency: memberDependency,
		AdminDependency:  adminDependency,
		Dependency:       dependency,

		JwtKey: jwtKey,
		Engine: engine,
	}
	router := apps.NewRouter(routerParent)

	// Run Server
	if port == "" {
		port = ":3001"
	}
	err = router.Engine.Run(port)
	helpers.FatalError(err)
}
