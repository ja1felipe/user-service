package models

import "gorm.io/gorm"

type CreditCard struct {
	BaseModel
	Number     string `json:"number" gorm:"not null; unique"`
	CVC        string `json:"cvc" gorm:"not null;"`
	Validate   string `json:"validate" gorm:"not null"`
	CCFullname string `json:"ccfullname" gorm:"not null"`
	UserID     uint   `json:"userId" gorm:"not null;"`
	User       User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c *CreditCard) BeforeCreate(tx *gorm.DB) (err error) {
	return CreditCardValidation(c)
}
