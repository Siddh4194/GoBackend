package userController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

// TODO: add the validation of the mobile number and the email
func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	userId, _ := strconv.Atoi(id) // make sure to handle error in real app

	if err := c.ShouldBindJSON(&user); err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	fields := map[string]interface{}{
		"Name":     user.Name,
		"Phone":    user.Phone,
		"Email":    user.Email,
		"Password": user.Password,
		"Address":  user.Address,
		// exclude CreatedAt, CartItems, etc. if not to be updated
	}

	if err := userRepo.UpdateFieldsByID(userId, fields); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}
	response.SendRespond(c, http.StatusOK, "user updated successfully")
}
