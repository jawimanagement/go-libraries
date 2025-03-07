package models

import (
	"time"

	"gorm.io/gorm"
)

type LanguageTableModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	TableTarget NullString `gorm:"column:table_target" json:"table_target"`
	PrimaryKey  NullString `gorm:"column:primary_key" json:"primary_key"`
	ColumnName  NullString `gorm:"column:column_name" json:"column_name"`
	Language    NullString `gorm:"column:language" json:"language"`
	Value       NullString `gorm:"column:value" json:"value"`
	CreatedFields
	UpdatedFields
}

type LanguageTableModelResponse = LanguageTableModel

func (p *LanguageTableModel) TableName() string {
	return "peserta"
}

func (p *LanguageTableModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *LanguageTableModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
