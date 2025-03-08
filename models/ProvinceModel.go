package models

import "gorm.io/gorm"

type ProvinceModel struct {
	ProvID     NullString `gorm:"column:prov_id;primary_key" json:"prov_id"`
	ProvName   NullString `gorm:"column:prov_name" json:"prov_name"`
	LocationID NullString `gorm:"column:locationid" json:"locationid"`
	Status     int        `gorm:"column:status;default:1" json:"status"`
}

func (p *ProvinceModel) TableName() string {
	return "provinces"
}

func (p *ProvinceModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *ProvinceModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
