package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Transaction struct {
	GormModel
	ProductID  uint `json:"product_id" form:"product_id"`
	UserID     uint `jsom:"user_id" form:"user_id"`
	Quantity   int  `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	TotalPrice int  `gorm:"not null" json:"total_price" form:"total_price" valid:"required~Total price is required"`
	Product    *Product
	User       *User
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (t *Transaction) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(t)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
