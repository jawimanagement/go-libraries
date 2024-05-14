package models

import "gorm.io/gorm"

type JadwalTgrModel struct {
	ID              string     `gorm:"column:id;primary_key" json:"id"`
	Partai          NullString `gorm:"column:partai" json:"partai"`
	Arena           int8       `gorm:"column:arena" json:"arena"`
	Corner          int8       `gorm:"column:corner" json:"corner"`
	Cualification   int8       `gorm:"column:cualification" json:"cualification"`
	Contingent      NullString `gorm:"column:contingent" json:"contingent"`
	EventID         NullString `gorm:"column:event_id" json:"event_id"`
	KategoriID      NullString `gorm:"column:kategori_id" json:"kategori_id"`
	GolonganID      NullString `gorm:"column:golongan_id" json:"golongan_id"`
	GenderID        NullString `gorm:"column:gender_id" json:"gender_id"`
	PesertaUniqueID NullString `gorm:"column:peserta_unique_id" json:"peserta_unique_id"`
}

func (p *JadwalTgrModel) TableName() string {
	return "temp_jadwal_tgr"
}

func (p *JadwalTgrModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTgrModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
