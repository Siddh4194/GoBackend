package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

// TODO: add the validation of the mobile number and the email
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Call the repository method to create a new user
	if err := userRepo.Insert(&user); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response.SendRespond(c, http.StatusCreated, user)
}