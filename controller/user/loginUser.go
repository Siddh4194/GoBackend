package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}
}