package server

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rumiani/rate-api/internal/middleware"
	"github.com/rumiani/rate-api/internal/server/routes"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.RateLimiter(5, 10))

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigin},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api/v1")

	routes.RegisterAssetRoutes(api.Group("/assets"))
	return r
}
