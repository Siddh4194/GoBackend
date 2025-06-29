package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/service"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	service.InitilizeAllRepo(db)
	api := r.Group("/api")
	{
		UserRoutes(api)
		OrderRoutes(api)
		FeedRoutes(api)
		ProductRoutes(api)
	}

	return r
}
