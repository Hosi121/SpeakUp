package main

import (
	"context"
	"log"

	"github.com/Hosi121/SpeakUp/config"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/middlewares"
	"github.com/Hosi121/SpeakUp/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get DSN and set up database connection
	dsn := config.GetDSN()

	apiKey := config.GetOpenAIKey()

	// Initialize gin
	r := gin.Default()
	r.Static("/upload", "./backend/upload")
	r.Use(middlewares.CORSMiddleware())

	// Ent client creation
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}
	defer client.Close()

	// Migrate the database schema
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Unauthenticated routes
	routes.SupabaseAuthRoutes(r, client)

	// Create a group for protected routes
	protected := r.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())

	routes.ChatRoute(r, apiKey)

	// Protected routes with Ent client
	routes.ProtectedRoutes(protected, client)

	// Run the server on port 8081
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
