package FeedController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func GetFeed(c *gin.Context) {
	var categories []models.Category
	if err := DB.Preload("Products").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}