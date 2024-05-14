package models

import (
	"time"

	"gorm.io/gorm"
)

type JadwalTandingModel struct {
	ID             string     `gorm:"column:id;primary_key" json:"id"`
	EventID        NullString `gorm:"column:event_id" json:"event_id"`
	Partai         NullString `gorm:"column:partai" json:"partai"`
	Arena          int        `gorm:"column:arena;default:null" json:"arena"`
	RedCorner      NullString `gorm:"column:red_corner" json:"red_corner"`
	BlueCorner     NullString `gorm:"column:blue_corner" json:"blue_corner"`
	ContingentRed  NullString `gorm:"column:contingent_red" json:"contingent_red"`
	ContingentBlue NullString `gorm:"column:contingent_blue" json:"contingent_blue"`
	KelasID        NullString `gorm:"column:kelas_id" json:"kelas_id"`
	GolonganID     NullString `gorm:"column:golongan_id" json:"golongan_id"`
	GenderID       NullString `gorm:"column:gender_id" json:"gender_id"`
	Status         int        `gorm:"column:status;default:0" json:"status"`
	Score          NullString `gorm:"column:score" json:"score"`
	WinBy          int        `gorm:"column:win_by;default:null" json:"win_by"`
	Winner         int        `gorm:"column:winner;default:null" json:"winner"`
	Round          int        `gorm:"column:round;default:0" json:"round"`
	Cualification  int        `gorm:"column:cualification;default:null" json:"cualification"`
	Inc            int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy      NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy      NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *JadwalTandingModel) TableName() string {
	return "temp_jadwal_tanding"
}

func (p *JadwalTandingModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTandingModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTandingModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JadwalTandingModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
