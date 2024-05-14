package models

import (
	"time"

	"gorm.io/gorm"
)

type BakuJudgeScoreModel struct {
	ID           string     `gorm:"column:id;primary_key" json:"id"`
	JadwalBakuID NullString `gorm:"column:jadwal_baku_id" json:"jadwal_baku_id"`
	JuriID       NullString `gorm:"column:juri_id" json:"juri_id"`
	Jurus1       NullString `gorm:"column:jurus1;default:0" json:"jurus1"`
	Jurus2       NullString `gorm:"column:jurus2;default:0" json:"jurus2"`
	Jurus3       NullString `gorm:"column:jurus3;default:0" json:"jurus3"`
	Jurus4       NullString `gorm:"column:jurus4;default:0" json:"jurus4"`
	Jurus5       NullString `gorm:"column:jurus5;default:0" json:"jurus5"`
	Jurus6       NullString `gorm:"column:jurus6;default:0" json:"jurus6"`
	Jurus7       NullString `gorm:"column:jurus7;default:0" json:"jurus7"`
	Jurus8       NullString `gorm:"column:jurus8;default:0" json:"jurus8"`
	Jurus9       NullString `gorm:"column:jurus9;default:0" json:"jurus9"`
	Jurus10      NullString `gorm:"column:jurus10;default:0" json:"jurus10"`
	Jurus11      NullString `gorm:"column:jurus11;default:0" json:"jurus11"`
	Agility      NullString `gorm:"column:agility;default:0" json:"agility"`
	TimePenalty  NullString `gorm:"column:time_penalty" json:"time_penalty"`
	DressPenalty NullString `gorm:"column:dress_penalty" json:"dress_penalty"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy    NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *BakuJudgeScoreModel) TableName() string {
	return "temp_baku_judge_score"
}

func (p *BakuJudgeScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *BakuJudgeScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *BakuJudgeScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *BakuJudgeScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
