package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"spt/docs"
	"spt/internal/db"
	"spt/internal/router"
	sptLog "spt/internal/util"
)

// @host     127.0.0.1:{{.Port}}
// @BasePath /api/v1
func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Setup log
	sptLog.SetupLog()

	// Create a database service instance
	dbService := db.Instance()

	// Create a Gin router
	router := SetupRouter(dbService)

	// Start the server on PORT
	router.Run(":" + os.Getenv("PORT"))
}

func SetupRouter(dbService db.Service) *gin.Engine {
	log.Infoln("SetupRouter()...")
	router := router.InitRouter()
	// Set up swagger info
	docs.SwaggerInfo.Title = "spt API"
	docs.SwaggerInfo.Description = "This is spt api server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/ping", func(c *gin.Context) {
		health := dbService.Health()
		c.JSON(http.StatusOK, health)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
