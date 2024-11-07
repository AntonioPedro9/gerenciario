package main

import (
	"server/cmd/api/routes"
	"server/internals/database"
	"server/internals/repositories"
	"server/pkg/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func init() {
	logs.InitLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// database instance
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// run migrations
	database.RunMigrations(db)

	// gin configurations
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"localhost"})

	// cors configurations
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// repositories
	userRepository := repositories.NewUserRepository(db)
	customerRepository := repositories.NewCustomerRepository(db)
	jobRepository := repositories.NewJobRepository(db)
	budgetRepository := repositories.NewBudgetRepository(db)

	// routes
	routes.SetupAuthRoutes(r, userRepository)
	routes.SetupUserRoutes(r, userRepository)
	routes.SetupCustomerRoutes(r, customerRepository, userRepository)
	routes.SetupJobRoutes(r, jobRepository, userRepository)
	routes.SetupBudgetRoutes(r, budgetRepository, userRepository)

	r.Run()
}
