package router

import (
	"github.com/gin-gonic/gin"
	productController "github.com/siddhant/shetkariBasketGo/controller/product"
)

func ProductRoutes(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.GET("/", productController.GetAll)
	}
}