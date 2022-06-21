package models

type CreditCard struct {
	BaseModel
	Number     string `json:"number" gorm:"not null; unique"`
	CVC        string `json:"cvc" gorm:"not null;"`
	Validation string `json:"validation" gorm:"not null"`
	CCFullname string `json:"ccfullname" gorm:"not null"`
	UserID     uint   `json:"userId" gorm:"not null;"`
}
