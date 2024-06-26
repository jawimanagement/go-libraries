package models

import (
	"time"

	"gorm.io/gorm"
)

type TgrJudgeScoreModel struct {
	ID        string     `gorm:"column:id;primary_key" json:"id"`
	JadwalID  NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	GroupID   NullString `gorm:"column:group_id" json:"group_id"`
	Jurus1    NullString `gorm:"column:jurus1;default:0" json:"jurus1"`
	Jurus2    NullString `gorm:"column:jurus2;default:0" json:"jurus2"`
	Jurus3    NullString `gorm:"column:jurus3;default:0" json:"jurus3"`
	Jurus4    NullString `gorm:"column:jurus4;default:0" json:"jurus4"`
	Jurus5    NullString `gorm:"column:jurus5;default:0" json:"jurus5"`
	Jurus6    NullString `gorm:"column:jurus6;default:0" json:"jurus6"`
	Jurus7    NullString `gorm:"column:jurus7;default:0" json:"jurus7"`
	Jurus8    NullString `gorm:"column:jurus8;default:0" json:"jurus8"`
	Jurus9    NullString `gorm:"column:jurus9;default:0" json:"jurus9"`
	Jurus10   NullString `gorm:"column:jurus10;default:0" json:"jurus10"`
	Jurus11   NullString `gorm:"column:jurus11;default:0" json:"jurus11"`
	Jurus12   NullString `gorm:"column:jurus12;default:0" json:"jurus12"`
	Jurus13   NullString `gorm:"column:jurus13;default:0" json:"jurus13"`
	Jurus14   NullString `gorm:"column:jurus14;default:0" json:"jurus14"`
	Agility   NullString `gorm:"column:agility;default:0" json:"agility"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *TgrJudgeScoreModel) TableName() string {
	return "tgr_judge_score"
}

func (p *TgrJudgeScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrJudgeScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrJudgeScoreModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TgrJudgeScoreModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
