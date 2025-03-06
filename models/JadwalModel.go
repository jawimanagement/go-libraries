package models

import (
	"time"

	"gorm.io/gorm"
)

type JadwalModel struct {
	ID             NullString `gorm:"column:id;primary_key" json:"id"`
	Partai         int64      `gorm:"column:partai" json:"partai"`
	EventID        NullString `gorm:"column:event_id" json:"event_id"`
	KategoriID     NullString `gorm:"column:kategori_id" json:"kategori_id"`
	GolonganID     NullString `gorm:"column:golongan_id" json:"golongan_id"`
	GenderID       NullString `gorm:"column:gender_id" json:"gender_id"`
	KelasTandingID NullString `gorm:"column:kelas_id" json:"kelas_id"`
	SchemaNumber   int64      `gorm:"column:schema_number" json:"schema_number"`
	Round          int        `gorm:"column:round;default:1" json:"round"`
	Cualification  int        `gorm:"column:cualification;default:1" json:"cualification"`
	Arena          int        `gorm:"column:arena" json:"arena"`
	Date           NullDate   `gorm:"column:date" json:"date"`
	Sesi           int        `gorm:"column:sesi" json:"sesi"`
	Pool           int16      `gorm:"column:pool;default:1" json:"pool"`
	RedCorner      NullString `gorm:"column:red_corner" json:"red_corner"`
	BlueCorner     NullString `gorm:"column:blue_corner" json:"blue_corner"`
	RedBye         int16      `gorm:"column:red_bye" json:"red_bye"`
	BlueBye        int16      `gorm:"column:blue_bye" json:"blue_bye"`
	ScheduleLevel  int16      `gorm:"column:schedule_level" json:"schedule_level"`
	Score          NullString `gorm:"column:score" json:"score"`
	Status         int        `gorm:"column:status" json:"status"`
	ScheduleType   int        `gorm:"column:schedule_type;default:1" json:"schedule_type"`
	IsReleased     int        `gorm:"column:is_released;default:0" json:"is_released"`
	Winner         NullString `gorm:"column:winner" json:"winner"`
	WinBy          int        `gorm:"column:win_by" json:"win_by"`
	INC            int        `gorm:"column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

func (p *JadwalModel) TableName() string {
	return "jadwal"
}

func (p *JadwalModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *JadwalModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}

func (p *JadwalModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
