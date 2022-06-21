package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

func (bh BaseHandler) Get(c *gin.Context) {

	var user models.User

	if result := bh.DB.Where("id = ?", c.Param("id")).First(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
