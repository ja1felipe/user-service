package v1routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) {
	main := router.Group("api/v1")
	UserRoutes(main, db)
}
