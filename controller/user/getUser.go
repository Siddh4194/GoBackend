package userController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func GetAllUsers(c *gin.Context) {
	users , err := userRepo.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid user ID", err)
		return
	}
	
	users , err := userRepo.GetWithCartDataById(userID)
	
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Failed to fetch the user data", err)
		return
	}
	
	response.SendRespond(c,http.StatusOK, users)
}