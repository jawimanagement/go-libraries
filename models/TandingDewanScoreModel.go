package models

import (
	"time"

	"gorm.io/gorm"
)

type TandingDewanScoreModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID  NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	Corner    int        `gorm:"column:corner" json:"corner"`
	Score     NullString `gorm:"column:score" json:"score"`
	Type      int        `gorm:"column:type" json:"type"`
	Status    int        `gorm:"column:status;default:0" json:"status"`
	Round     int        `gorm:"column:round" json:"round"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *TandingDewanScoreModel) TableName() string {
	return "tanding_dewan_score"
}

func (p *TandingDewanScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingDewanScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingDewanScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingDewanScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
