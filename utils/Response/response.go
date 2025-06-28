package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string, err error) {
	if err != nil {
		c.JSON(status, gin.H{
			"error":true,
			"response":   message,
			"details": err.Error(),
		})
	} else {
		c.JSON(status, gin.H{
			"error":true,
			"response": message,
		})
	}
}

func SendRespond(c *gin.Context, status int, data interface{}) {
	if data == nil {
		c.Status(status)
		return
	}
	encryptedData, err := respondEncrypted(data)
	if err != nil {
		RespondError(c,http.StatusInternalServerError, "Failed to encrypt data", err)
	}
	c.JSON(status, gin.H{
		"error":   false,
		"response": encryptedData,
	})
}