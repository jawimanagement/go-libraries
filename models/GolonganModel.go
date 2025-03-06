package models

import "gorm.io/gorm"

type GolonganModel struct {
	ID     NullString `gorm:"column:id;primary_key" json:"id"`
	Name   NullString `gorm:"column:name" json:"name"`
	MinAge NullInt64  `gorm:"column:min_age" json:"min_age"`
	MaxAge NullInt64  `gorm:"column:max_age" json:"max_age"`
	Active int        `gorm:"column:active;default:1" json:"active"`
}

type GolonganModelResponse struct {
	ID     NullString `gorm:"column:id;primary_key" json:"id"`
	Name   NullString `gorm:"column:name" json:"name"`
	MinAge NullInt64  `gorm:"column:min_age" json:"min_age"`
	MaxAge NullInt64  `gorm:"column:max_age" json:"max_age"`
	Active int        `gorm:"column:active;default:1" json:"active"`
}

func (p *GolonganModel) TableName() string {
	return "golongan"
}

func (p *GolonganModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GolonganModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
