package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// InitRouter initializes the Gin engine with necessary middleware and routes.
func InitRouter() *gin.Engine {
	// setup gin
	r := setupGin()

	// setup routes
	setupRoutes(r)

	return r
}

// setupGin configures the Gin engine with necessary middleware.
func setupGin() *gin.Engine {
	gin.SetMode(os.Getenv("RUN_MODE"))
	gin.ForceConsoleColor()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	return r
}

func setupRoutes(r *gin.Engine) {
	// Create a new route group for API version 1
	rg := r.Group("/api/v1")
	{
		// Register project-related routes
		RegisterProjectRoutes(rg)
	}
}
