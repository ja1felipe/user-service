package models

type Address struct {
	BaseModel
	UserID     uint   `json:"userId" gorm:"not null;"`
	Street     string `json:"street" gorm:"not null;"`
	ZipCode    uint   `json:"zipcode" gorm:"not null;"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
}
