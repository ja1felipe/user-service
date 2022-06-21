package migrations

import (
	"github.com/ja1felipe/user-service/pkg/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.BaseModel{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.CreditCard{})
}
