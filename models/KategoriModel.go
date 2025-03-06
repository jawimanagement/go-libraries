package models

import "gorm.io/gorm"

type KategoriModel struct {
	ID         NullString `gorm:"column:id;primary_key" json:"id"`
	Name       NullString `gorm:"column:name" json:"name"`
	MinAthlete NullInt64  `gorm:"column:min_athlete" json:"min_athlete"`
	MaxAthlete NullInt64  `gorm:"column:max_athlete" json:"max_athlete"`
	Active     int        `gorm:"column:active;default:1" json:"active"`
}

type KategoriModelResponse = KategoriModel

func (p *KategoriModel) TableName() string {
	return "kategori"
}

func (p *KategoriModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *KategoriModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
