package categoryController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func Update(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	categoryId, _ := strconv.Atoi(id) // make sure to handle error in real app

	if err := c.ShouldBindJSON(&category); err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	fields := map[string]interface{}{
		"Name":     category.Name,
		"Phone":    category.Description,
	}

	if err := categoryRepo.UpdateFieldsByID(categoryId, fields); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}
	response.SendRespond(c, http.StatusOK, "category updated successfully")
}
