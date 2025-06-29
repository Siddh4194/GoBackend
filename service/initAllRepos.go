package service

import (
	repository "github.com/siddhant/shetkariBasketGo/Repository"
	categoryController "github.com/siddhant/shetkariBasketGo/controller/Category"
	orderController "github.com/siddhant/shetkariBasketGo/controller/order"
	productController "github.com/siddhant/shetkariBasketGo/controller/product"
	userController "github.com/siddhant/shetkariBasketGo/controller/user"
	"gorm.io/gorm"
)

func InitilizeAllRepo(db *gorm.DB) {

	// user repository
	// This is where we initialize the user repository and pass it to the user controller
	repo := repository.NewGormUserRepository(db)
	userController.InitUserController(repo)

	// category repository initilization with the controller
	categoryRepo := repository.NewGormCategoryRepository(db)
	categoryController.InitCategoryController(categoryRepo)

	// order repository initilization with the controller
	orderRepo := repository.NewGormOrderRepository(db)
	orderController.InitOrderController(orderRepo)

	// product repository initilization with the controller
	productRepo := repository.NewGormProductRepository(db)
	productController.InitProductController(productRepo)
}