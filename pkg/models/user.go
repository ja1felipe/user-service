package models

import (
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FullName       string     `json:"fullname"`
	Email          string     `json:"email" gorm:"not null; unique"`
	Role           string     `json:"role" gorm:"default:user;"`
	Identification string     `json:"identification" gorm:"unique"`
	Address        Address    `gorm:"foreignKey:UserID;"`
	CreditCard     CreditCard `gorm:"foreignKey:UserID;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return BasicUserValidation(u)
}
