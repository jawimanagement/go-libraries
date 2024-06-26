package models

import (
	"time"

	"gorm.io/gorm"
)

type GandaJudgeScoreModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID   NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	GroupID    NullString `gorm:"column:group_id" json:"group_id"`
	SerangBela NullString `gorm:"column:serang_bela;default:0" json:"serang_bela"`
	Kompak     NullString `gorm:"column:kompak;default:0" json:"kompak"`
	Hayat      NullString `gorm:"column:hayat;default:0" json:"hayat"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy  NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy  NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *GandaJudgeScoreModel) TableName() string {
	return "ganda_dewan_score"
}

func (p *GandaJudgeScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaJudgeScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaJudgeScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *GandaJudgeScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
