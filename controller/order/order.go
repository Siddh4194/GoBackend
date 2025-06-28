package orderController

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

func GetAllOrder(c *gin.Context) {
	var orders []models.Order
	if err := DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
