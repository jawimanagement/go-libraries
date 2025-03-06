package models

import (
	"time"

	"gorm.io/gorm"
)

type EventDetailModel struct {
	ID         NullString `gorm:"column:id;primary_key" json:"id"`
	EventID    NullString `gorm:"column:event_id" json:"event_id"`
	KategoriID NullString `gorm:"column:kategori_id" json:"kategori_id"`
	GenderID   NullString `gorm:"column:gender_id" json:"gender_id"`
	GolonganID NullString `gorm:"column:golongan_id" json:"golongan_id"`
	KelasID    NullString `gorm:"column:kelas_id" json:"kelas_id"`
	Min        NullInt64  `gorm:"column:min;default:3" json:"min"`
	Sec        NullInt64  `gorm:"column:sec;default:0" json:"sec"`
	Prices     float64    `gorm:"column:prices;default:0" json:"prices"`
	MaxBabak   int        `gorm:"column:max_babak;default:3" json:"max_babak"`
	INC        int        `gorm:"<-:false;column:inc" json:"inc"`
	Active     int        `gorm:"column:active;default:1" json:"active"`
	CreatedFields
	UpdatedFields
}

type EventDetailModelResponse struct {
	EventDetailModel
	EventInfo    EventsModelResponse       `gorm:"->;foreignKey:EventID;references:ID" json:"event_info"`
	KategoriInfo KategoriModelResponse     `gorm:"->;foreignKey:KategoriID;references:ID" json:"kategori_info"`
	GenderInfo   GenderModelResponse       `gorm:"->;foreignKey:GenderID;references:ID" json:"gender_info"`
	GolonganInfo GolonganModelResponse     `gorm:"->;foreignKey:GolonganID;references:ID" json:"golongan_info"`
	KelasInfo    KelasTandingModelResponse `gorm:"->;foreignKey:KelasID;references:ID" json:"kelas_info"`
}

func (p *EventDetailModel) TableName() string {
	return "event_detail"
}

func (p *EventDetailModelResponse) TableName() string {
	return "event_detail"
}

func (p *EventDetailModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventDetailModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
