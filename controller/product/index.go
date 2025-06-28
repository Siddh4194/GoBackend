package productController

import repository "github.com/siddhant/shetkariBasketGo/Repository"

var productRepo repository.ProductRepository // Interface

func InitProductController(repo repository.ProductRepository) {
	productRepo = repo
}
