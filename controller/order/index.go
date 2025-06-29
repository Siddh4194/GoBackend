package orderController

import repository "github.com/siddhant/shetkariBasketGo/Repository"

var orderRepo repository.OrderRepository

func InitOrderController(repo repository.OrderRepository) {
	orderRepo = repo
}