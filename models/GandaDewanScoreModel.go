package models

import (
	"time"

	"gorm.io/gorm"
)

type GandaDewanScoreModel struct {
	ID       string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	GroupID  NullString `gorm:"column:group_id" json:"group_id"`
	Time     NullString `gorm:"column:time;default:0" json:"time"`
	Drop     NullString `gorm:"column:drop;default:0" json:"drop"`
	FallOut  NullString `gorm:"column:fall_out;default:0" json:"fall_out"`
	Hold     NullString `gorm:"column:hold;default:0" json:"hold"`
	OutArena NullString `gorm:"column:out_arena;default:0" json:"out_arena"`
	Dress    NullString `gorm:"column:dress;default:0" json:"dress"`
	Voice    NullString `gorm:"column:voice;default:0" json:"voice"`
	CreatedFields
	UpdatedFields
}

func (p *GandaDewanScoreModel) TableName() string {
	return "ganda_dewan_score"
}

func (p *GandaDewanScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *GandaDewanScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}

func (p *GandaDewanScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaDewanScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
