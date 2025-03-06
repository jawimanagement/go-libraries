package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name" json:"name"`
	Price       float64    `gorm:"column:price" json:"price"`
	Description NullString `gorm:"column:description" json:"description"`
	Type        int        `gorm:"column:type;default:1"`
	Volume      int        `gorm:"column:volume;default:1" json:"volume"`
	Status      int        `gorm:"column:status;default:1" json:"status"`
	INC         int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type ProductModelResponse = ProductModel

func (p *ProductModel) TableName() string {
	return "product"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *ProductModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
