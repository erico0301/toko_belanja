package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required~Product Title is required"`
	Price       int    `gorm:"not null" json:"price" form:"price" valid:"required~Product Price is required, range(0|50000000)~Product Price is out of range"`
	Stock       int    `gorm:"not null" json:"stock" form:"stock" valid:"required~Stok is required, range(5|999999)~minimum stock is 5"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Category    *Category
	Transaction []Transaction `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"transactions"`
}

type GetProduct struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	Price      int        `json:"price"`
	Stock      int        `json:"stock"`
	CategoryID uint       `json:"category_id"`
	CreatedAt  *time.Time `json:"created_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
