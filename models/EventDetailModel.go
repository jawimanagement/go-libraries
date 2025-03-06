package models

import (
	"time"

	"gorm.io/gorm"
)

type EventDetailModel struct {
	ID          		NullString	`gorm:"column:id;primary_key" json:"id"`
	EventID				NullString	`gorm:"column:event_id" json:"event_id"`
	KategoriID			NullString	`gorm:"column:kategori_id" json:"kategori_id"`
	GenderID			NullString	`gorm:"column:gender_id" json:"gender_id"`
	GolonganID			NullString	`gorm:"column:golongan_id" json:"golongan_id"`
	KelasID				NullString	`gorm:"column:kelas_id" json:"kelas_id"`
	Min					NullInt64	`gorm:"column:min;default:3" json:"min"`
	Sec					NullInt64	`gorm:"column:sec;default:0" json:"sec"`
	Prices				float64		`gorm:"column:prices;default:0" json:"prices"`
	MaxBabak			int			`gorm:"column:max_babak;default:3" json:"max_babak"`
	INC					int			`gorm:"<-:false;column:inc" json:"inc"`
	Active				int			`gorm:"column:active;default:1" json:"active"`
	CreatedAt   		*time.Time 	`gorm:"column:created_at" json:"created_at"`
	CreatedBy   		NullString 	`gorm:"column:created_by" json:"created_by"`
	UpdatedAt   		*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   		NullString 	`gorm:"column:updated_by" json:"updated_by"`
}

type EventDetailModelResponse struct {
	ID          		NullString	`gorm:"column:id;primary_key" json:"id"`
	EventID				NullString	`gorm:"column:event_id" json:"event_id"`
	KategoriID			NullString	`gorm:"column:kategori_id" json:"kategori_id"`
	GenderID			NullString	`gorm:"column:gender_id" json:"gender_id"`
	GolonganID			NullString	`gorm:"column:golongan_id" json:"golongan_id"`
	KelasID				NullString	`gorm:"column:kelas_id" json:"kelas_id"`
	Min					NullInt64	`gorm:"column:min;default:3" json:"min"`
	Sec					NullInt64	`gorm:"column:sec;default:0" json:"sec"`
	Prices				float64		`gorm:"column:prices;default:0" json:"prices"`
	MaxBabak			int			`gorm:"column:max_babak;default:3" json:"max_babak"`
	INC					int			`gorm:"<-:false;column:inc" json:"inc"`
	Active				int			`gorm:"column:active;default:1" json:"active"`
	EventInfo			EventsModelResponse	`gorm:"<-;foreignKey:ID;references:EventID" json:"event_info"`
	KategoriInfo		KategoriModelResponse	`gorm:"<-;foreignKey:ID;references:KategoriID" json:"kategori_info"`
	GenderInfo			GenderModelResponse	`gorm:"<-;foreignKey:ID;references:GenderID" json:"gender_info"`
	GolonganInfo		GolonganModelResponse	`gorm:"<-;foreignKey:ID;references:GolonganID" json:"golongan_info"`
	KelasInfo			KelasTandingModelResponse	`gorm:"<-;foreignKey:ID;references:KelasID" json:"kelas_info"`
	CreatedAt   		*time.Time 	`gorm:"column:created_at" json:"created_at"`
	CreatedBy   		NullString 	`gorm:"column:created_by" json:"created_by"`
	UpdatedAt   		*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   		NullString 	`gorm:"column:updated_by" json:"updated_by"`
}


func (p *EventDetailModel) TableName() string {
	return "event_detail"
}

func (p *EventDetailModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *EventDetailModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}