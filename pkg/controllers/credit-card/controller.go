package creditcard

import "gorm.io/gorm"

type BaseHandler struct {
	DB *gorm.DB
}

func NewCreditCardHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{DB: db}
}
