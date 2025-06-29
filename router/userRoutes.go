package router

import (
	"github.com/gin-gonic/gin"
	userController "github.com/siddhant/shetkariBasketGo/controller/user"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/", userController.GetAllUsers)
		users.GET("/:id", userController.GetUserById)
	}
}