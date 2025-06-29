package categoryController

import (
	repository "github.com/siddhant/shetkariBasketGo/Repository"
)

var categoryRepo repository.CategoryRepository

func InitCategoryController(repo repository.CategoryRepository) {
	categoryRepo = repo
}