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
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy   NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   NullString `gorm:"column:updated_by" json:"updated_by"`
}

type ProductPackageModelResponse = ProductPackageModel

func (p *ProductPackageModel) TableName() string {
	return "product_package"
}

func (p *ProductPackageModel) BeforeCreate(tx *gorm.DB) (err *error) {
	return
}

func (p *ProductPackageModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	return
}
