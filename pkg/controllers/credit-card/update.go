package creditcard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

type updateBodyRequest struct {
	Number     string `json:"number"`
	CVC        string `json:"cvc"`
	Validate   string `json:"validate"`
	CCFullname string `json:"ccfullname"`
}

func (bh BaseHandler) Update(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	if err := bh.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot find user: " + err.Error(),
		})
		return
	}

	body := createBodyRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var creditCard models.CreditCard

	if err := bh.DB.Where("UserId = ?", userId).First(&creditCard).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot find credit card: " + err.Error(),
		})
		return
	}

	updatedCreditCard := models.CreditCard{Number: body.Number, CVC: body.CVC, CCFullname: body.CCFullname, Validate: body.Validate}
	err := bh.DB.Model(&creditCard).Updates(&updatedCreditCard).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "cannot update credit card: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, creditCard)
}
