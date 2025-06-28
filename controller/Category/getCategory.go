package categoryController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func GetAll(c *gin.Context) {
	categories, err := categoryRepo.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to fetch categories", err)
	}
	response.SendRespond(c, http.StatusOK, categories)
}

func GetAllWithProducts(c *gin.Context){
	categories, err := categoryRepo.GetAllWithProduct()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to fetch categories", err)
	}
	response.SendRespond(c, http.StatusOK, categories)
}