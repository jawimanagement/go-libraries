package models

import (
	"time"

	"gorm.io/gorm"
)

type TimeKeeperModel struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	EventID    NullString `gorm:"event_id" json:"event_id"`
	KategoriID NullString `gorm:"kategori_id" json:"kategori_id"`
	JadwalID   NullString `gorm:"column:jadwal_id" json:"jadwal_id"`
	Arena      int        `gorm:"column:arena" json:"arena"`
	Round      int        `gorm:"column:round;default:0" json:"round"`
	Status     int        `gorm:"column:status;default:0" json:"status"`
	StartAt    *time.Time `gorm:"column:start_at" json:"start_at"`
	StopAt     *time.Time `gorm:"column:stop_at" json:"stop_at"`
	PausedAt   string     `gorm:"column:paused_at" json:"paused_at"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy  NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy  NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *TimeKeeperModel) TableName() string {
	return "temp_time_keeper"
}

func (p *TimeKeeperModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TimeKeeperModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *TimeKeeperModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *TimeKeeperModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
