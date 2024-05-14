package models

import (
	"time"

	"gorm.io/gorm"
)

type JadwalTgrPesertaModel struct {
	ID              string     `gorm:"column:id;primary_key" json:"id"`
	Name            NullString `gorm:"column:name" json:"name"`
	JadwalTgrID     NullString `gorm:"column:jadwal_tgr_id" json:"jadwal_tgr_id"`
	PesertaUniqueID NullString `gorm:"column:peserta_unique_id" json:"peserta_unique_id"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy       NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy       NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *JadwalTgrPesertaModel) TableName() string {
	return "temp_jadwal_tgr_peserta"
}

func (p *JadwalTgrPesertaModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrPesertaModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrPesertaModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrPesertaModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
