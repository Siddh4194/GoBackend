package categoryController

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddhant/shetkariBasketGo/models"
	response "github.com/siddhant/shetkariBasketGo/utils/Response"
)

func Post(c *gin.Context) {
	var Category models.Category
	data,err := c.GetRawData()

	if err != nil {
      c.AbortWithStatusJSON(400, "User is not defined")
      return
	}

   err = json.Unmarshal(data, &Category)
   if err != nil {
      c.AbortWithStatusJSON(400, "Bad Input")
      return
   }

	if err := categoryRepo.Insert(&Category); err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to add category", err)
		return
	}

	response.SendRespond(c, http.StatusCreated,nil)
}
