package userController

import (
	repository "github.com/siddhant/shetkariBasketGo/Repository"
)

var userRepo repository.UserRepository // Interface

func InitUserController(repo repository.UserRepository) {
	userRepo = repo
}