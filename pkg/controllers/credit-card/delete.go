package creditcard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

func (bh BaseHandler) Delete(c gin.Context) {
	var user models.User

	if err := bh.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	var creditCard models.CreditCard

	if err := bh.DB.Where("id = ?", user.ID).First(&creditCard).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "credit card not found",
		})
		return
	}

	err := bh.DB.Delete(&creditCard).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "error deleting credit card" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": "credit card deleted with success",
	})
}
