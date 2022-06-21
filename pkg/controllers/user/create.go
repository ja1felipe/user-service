package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

type createRequestBody struct {
	Email          string `json:"email"`
	FullName       string `json:"fullname"`
	Role           string `json:"role"`
	Identification string `json:"identification"`
}

func (bh BaseHandler) Create(c *gin.Context) {
	body := createRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User

	if bh.DB.Where("email = ?", body.Email).First(&user); user.Email != "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "email already registred",
		})
		return
	}

	user.Email = body.Email
	user.FullName = body.FullName
	user.Identification = body.Identification

	if body.Role != "" {
		user.Role = body.Role
	}

	err := bh.DB.Create(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "cannot create user " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}
