package models

import "gorm.io/gorm"

type CityModel struct {
	CityID   NullString `gorm:"column:city_id;primary_key" json:"city_id"`
	CityName NullString `gorm:"column:city_name" json:"city_name"`
	ProvID   NullString `gorm:"column:prov_id" json:"prov_id"`
}

func (p *CityModel) TableName() string {
	return "cities"
}

func (p *CityModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *CityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
