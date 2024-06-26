package models

import (
	"time"

	"gorm.io/gorm"
)

type TandingFinalScoreModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID  NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	Corner    int        `gorm:"column:corner" json:"corner"`
	Score     NullString `gorm:"column:score" json:"score"`
	Round     int        `gorm:"column:round" json:"round"`
	Status    int        `gorm:"column:status;default:0" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *TandingFinalScoreModel) TableName() string {
	return "tanding_final_score"
}

func (p *TandingFinalScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingFinalScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingFinalScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingFinalScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
