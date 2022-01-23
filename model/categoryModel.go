package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	GormModel
	Type              string    `gorm:"not null" json:"type" form:"type" valid:"required~Type of category is required"`
	SoldProductAmount int       `json:"sold_product_amount" form:"sold_product_amount"`
	Product           []Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"products"`
}

func (c *Category) BeforCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
