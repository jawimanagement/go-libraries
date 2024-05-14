package models

import (
	"time"

	"gorm.io/gorm"
)

type JadwalJurusBakuModel struct {
	ID            string     `gorm:"column:id;primary_key" json:"id"`
	Partai        NullString `gorm:"column:partai" json:"partai"`
	Arena         int        `gorm:"column:arena" json:"arena"`
	Corner        int        `gorm:"column:corner" json:"corner"`
	Cualification int        `gorm:"column:cualification" json:"cualification"`
	Contingent    NullString `gorm:"column:contingent" json:"contingent"`
	EventID       NullString `gorm:"column:event_id" json:"event_id"`
	KategoriID    NullString `gorm:"column:kategori_id" json:"kategori_id"`
	GenderID      NullString `gorm:"column:gender_id" json:"gender_id"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *JadwalJurusBakuModel) TableName() string {
	return "temp_jadwal_jurus_baku"
}

func (p *JadwalJurusBakuModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalJurusBakuModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalJurusBakuModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalJurusBakuModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
