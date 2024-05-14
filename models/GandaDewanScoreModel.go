package models

import (
	"time"

	"gorm.io/gorm"
)

type GandaDewanScoreModel struct {
	ID          string     `gorm:"column:id;primary_key" json:"id"`
	JadwalTgrID NullString `gorm:"column:jadwal_tgr_id" json:"jadwal_tgr_id"`
	Time        NullString `gorm:"column:time;default:0" json:"time"`
	Drop        NullString `gorm:"column:drop;default:0" json:"drop"`
	FallOut     NullString `gorm:"column:fall_out;default:0" json:"fall_out"`
	Hold        NullString `gorm:"column:hold;default:0" json:"hold"`
	OutArena    NullString `gorm:"column:out_arena;default:0" json:"out_arena"`
	Dress       NullString `gorm:"column:dress;default:0" json:"dress"`
	Voice       NullString `gorm:"column:voice;default:0" json:"voice"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy   NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *GandaDewanScoreModel) TableName() string {
	return "temp_ganda_dewan_score"
}

func (p *GandaDewanScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaDewanScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaDewanScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaDewanScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
