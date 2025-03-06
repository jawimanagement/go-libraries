package models

import (
	"time"

	"gorm.io/gorm"
)

type LogTableModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	TableTarget NullString `gorm:"column:table_target" json:"table_target"`
	PrimaryKey  NullString `gorm:"column:primary_key" json:"primary_key"`
	Value       NullString `gorm:"column:value" json:"value"`
	Status      int        `gorm:"column:status;default:1" json:"status"`
	INC         int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (p *LogTableModel) TableName() string {
	return "log_table"
}

func (p *LogTableModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *LogTableModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *LogTableModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *LogTableModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
