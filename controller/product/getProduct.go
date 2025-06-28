package productController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func GetAll(c *gin.Context) {
	products, err := productRepo.GetAll()
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to fetch categories", err)
	}
	response.SendRespond(c, http.StatusOK, products)
}