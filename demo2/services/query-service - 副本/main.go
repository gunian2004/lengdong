package main

import (
	"fmt"
	"log"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/middleware"
	"github.com/coldchain-traceability-system/database/mysql"
	"github.com/coldchain-traceability-system/database/redis"
	"github.com/coldchain-traceability-system/services/query-service/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig("../../")
	
	// Initialize database
	mysql.InitDB()
	
	// Initialize Redis
	redis.InitRedis()
	
	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.Mode)
	
	// Initialize Gin router
	r := gin.Default()
	
	// Apply middleware
	r.Use(middleware.CorsMiddleware())
	
	// Setup router
	router.SetupRouter(r)
	
	// Start server
	serverAddr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("Query service starting on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 