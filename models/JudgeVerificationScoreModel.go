package models

import (
	"time"

	"gorm.io/gorm"
)

type JudgeVerificationScoreModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID  NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	NilaiID   NullString `gorm:"column:nilai_id" json:"nilai_id"`
	Corner    int        `gorm:"column:corner" json:"corner"`
	Round     int        `gorm:"column:round" json:"round"`
	Status    int        `gorm:"column:status;default:0" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *JudgeVerificationScoreModel) TableName() string {
	return "judge_verification_score"
}

func (p *JudgeVerificationScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JudgeVerificationScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *JudgeVerificationScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *JudgeVerificationScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
