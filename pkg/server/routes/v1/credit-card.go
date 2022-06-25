package v1routes

import (
	"github.com/gin-gonic/gin"
	creditcard "github.com/ja1felipe/user-service/pkg/controllers/credit-card"
	"gorm.io/gorm"
)

func CreditCardRoutes(group *gin.RouterGroup, db *gorm.DB) {
	handler := creditcard.NewCreditCardHandler(db)
	cc := group.Group("cc")
	{
		cc.POST("/:id/", handler.Create)
	}
}
