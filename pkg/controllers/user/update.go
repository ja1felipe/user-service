package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

type updateRequestBody struct {
	Email          string `json:"email"`
	FullName       string `json:"fullname"`
	Role           string `json:"role"`
	Identification string `json:"identification"`
}

func (bh BaseHandler) Update(c *gin.Context) {
	body := updateRequestBody{}

	var user models.User

	if err := bh.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot find user: " + err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	updateUser := models.User{Email: body.Email, FullName: body.FullName, Role: body.Role, Identification: body.Identification}

	if err := models.BasicUserValidation(&updateUser); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "cannot update user: " + err.Error(),
		})
		return
	}

	err := bh.DB.Model(&user).Updates(&updateUser).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "cannot update user: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
