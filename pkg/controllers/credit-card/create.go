package creditcard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

type createBodyRequest struct {
	Number     string `json:"number"`
	CVC        string `json:"cvc"`
	Validate   string `json:"validate"`
	CCFullname string `json:"ccfullname"`
}

func (bh BaseHandler) Create(c *gin.Context) {
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

	creditCard.CVC = body.CVC
	creditCard.CCFullname = body.CCFullname
	creditCard.Number = body.Number
	creditCard.Validate = body.Validate
	creditCard.UserID = user.ID

	err := bh.DB.Create(&creditCard).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "cannot create credit card: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, creditCard)
}
