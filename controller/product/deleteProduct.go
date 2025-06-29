package productController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func Delete(c *gin.Context) {
	productId := c.Param("id")
	id, _ := strconv.Atoi(productId)
	if err := productRepo.Delete(id); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to delete the product", err)
	}
	response.SendRespond(c, http.StatusOK, "Product deleted successfully")
}