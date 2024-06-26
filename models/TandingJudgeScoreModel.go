package models

import (
	"time"

	"gorm.io/gorm"
)

type TandingJudgeScoreModel struct {
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

func (p *TandingJudgeScoreModel) TableName() string {
	return "temp_tanding_judge_score"
}

func (p *TandingJudgeScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingJudgeScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingJudgeScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TandingJudgeScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
