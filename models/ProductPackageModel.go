package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductPackageModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	Name        NullString `gorm:"column:name" json:"name"`
	Description NullString `gorm:"column:description" json:"description"`
	Status      int        `gorm:"column:status;default:1" json:"status"`
	INC         int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type ProductPackageModelResponse = ProductPackageModel

func (p *ProductPackageModel) TableName() string {
	return "product_package"
}

func (p *ProductPackageModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *ProductPackageModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
