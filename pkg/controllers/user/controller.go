package user

import "gorm.io/gorm"

type BaseHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{DB: db}
}
