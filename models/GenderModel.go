package models

import "gorm.io/gorm"

type GenderModel struct {
	ID    NullString `gorm:"column:id;primary_key" json:"id"`
	Name  NullString `gorm:"column:name" json:"name"`
	Alias NullString `gorm:"column:alias" json:"alias"`
}

type GenderModelResponse = GenderModel

func (p *GenderModel) TableName() string {
	return "gender"
}

func (p *GenderModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GenderModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
