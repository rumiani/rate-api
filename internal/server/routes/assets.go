package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rumiani/rate-api/internal/asset"
)

func RegisterAssetRoutes(rg *gin.RouterGroup) {
	// GET /assets → list all
	log.Println("here")
	rg.GET("/", func(c *gin.Context) {
		assets, err := asset.GetAllAssets()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch assets"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": assets})
	})

	// GET /assets/:code → single asset
	rg.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		a, err := asset.GetAssetByCode(code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "asset not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": a})
	})
}
