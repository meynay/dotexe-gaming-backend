package main

import (
	"log"
	"store/internal/delivery/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.SecurityHeaders())
	router.Use(middlewares.RateLimiter())
	router.Use(middlewares.CSRFMiddleware())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorizaion"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middlewares.APIKeyAuth())
	router.Use(gzip.Gzip(gzip.BestCompression))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "app works"})
	})
	public := router.Group("/public-api")
	private := router.Group("/private-api")
	private.Use(middlewares.AuthMiddleware())
	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi, public works"})
	})
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hi, private works"})
	})
	router.Run(":8080")
}
