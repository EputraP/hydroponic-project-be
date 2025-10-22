package main

import (
	"fmt"
	"hydroponic-be/internal/routes"
	"hydroponic-be/internal/util/logger"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := logger.Init("app.log"); err != nil {
		logger.Error("main", "Failed to initialize logger:", map[string]string{
			"error": err.Error(),
		})
		return
	}

	logger.Info("main", "Starting application...", nil)

	err := godotenv.Load()
	if err != nil {
		logger.Error("main", "Error loading .env file", map[string]string{
			"error": err.Error(),
		})
		return
	}

	srv := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("main",
		"Server is starting...", map[string]string{
			"port": port,
		})

	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Error("main", "Error running Gin server", map[string]string{
			"error": err.Error(),
		})
	}
}

func prepare() (handlers routes.Handlers, middlewares routes.Middlewares) {

	return
}
