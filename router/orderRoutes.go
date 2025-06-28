package router

import (
	"github.com/gin-gonic/gin"
	orderController "github.com/siddhant/shetkariBasketGo/controller/order"
)

func OrderRoutes(r *gin.RouterGroup) {
	orders := r.Group("/orders")
	{
		orders.GET("/", orderController.GetAllOrder)
	}
}