package v1routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ja1felipe/user-service/pkg/controllers/user"
	"gorm.io/gorm"
)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	handler := user.NewUserHandler(db)
	user := group.Group("user")
	{
		user.POST("/", handler.Create)
		user.PUT("/:id", handler.Update)
		user.DELETE("/:id", handler.Delete)
		user.GET("/:id", handler.Get)
	}
}
