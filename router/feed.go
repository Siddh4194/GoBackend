package router

import (
	"github.com/gin-gonic/gin"
	FeedController "github.com/siddhant/shetkariBasketGo/controller/Feed"
)

func FeedRoutes(r *gin.RouterGroup) {
	feedRoutes := r.Group("/feed")
	{
		feedRoutes.GET("/", FeedController.GetFeed)
	}
}