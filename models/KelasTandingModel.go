package models

import "gorm.io/gorm"

type KelasTandingModel struct {
	ID     NullString `gorm:"column:id;primary_key" json:"id"`
	Name   NullString `gorm:"column:name" json:"name"`
	Active int        `gorm:"column:active;default:1" json:"active"`
}

type KelasTandingModelResponse struct {
	ID     NullString `gorm:"column:id;primary_key" json:"id"`
	Name   NullString `gorm:"column:name" json:"name"`
	Active int        `gorm:"column:active" json:"active"`
}

func (p *KelasTandingModel) TableName() string {
	return "golongan"
}

func (p *KelasTandingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *KelasTandingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
