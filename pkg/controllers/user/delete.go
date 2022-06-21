package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/models"
)

func (bh BaseHandler) Delete(c *gin.Context) {

	var user models.User

	if result := bh.DB.Where("id = ?", c.Param("id")).First(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	err := bh.DB.Delete(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "error deleting user" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": "user deleted with success",
	})
}
