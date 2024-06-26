package models

import (
	"time"

	"gorm.io/gorm"
)

type TgrDewanScoreModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID  NullString `column:"jadwal_id" json:"jadwal_id"`
	GroupID   NullString `gorm:"column:group_id" json:"group_id"`
	Time      NullString `gorm:"column:time;default:0" json:"time"`
	OutArena  NullString `gorm:"column:out_arena;default:0" json:"out_arena"`
	Drop      NullString `gorm:"column:drop;default:0" json:"drop"`
	Dress     NullString `gorm:"column:dress;default:0" json:"dress"`
	Voice     NullString `gorm:"column:voice;default:0" json:"voice"`
	Hold      NullString `gorm:"column:hold;default:0" json:"hold"`
	Distance  NullString `gorm:"column:distance;default:0" json:"distance"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *TgrDewanScoreModel) TableName() string {
	return "tgr_dewan_score"
}

func (p *TgrDewanScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrDewanScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrDewanScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrDewanScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
