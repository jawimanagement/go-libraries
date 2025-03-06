package models

import (
	"time"

	"gorm.io/gorm"
)

type EventCoachModel struct {
	ID        NullString `gorm:"column:id;primary_key" json:"id"`
	EventID   NullString `gorm:"column:event_id" json:"event_id"`
	CoachID   NullString `gorm:"column:coach_id" json:"coach_id"`
	QRCode    NullString `gorm:"column:qrcode" json:"qrcode"`
	INC       int        `gorm:"<-:false;column:inc" json:"inc"`
	Status    int        `gorm:"column:status;default:1" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

type EventCoachModelResponse struct {
	EventCoachModel
	EventInfo EventsModelResponse `gorm:"->;foreignKey:EventID;references:ID" json:"event_info"`
	CoachInfo CoachModelResponse  `gorm:"->;foreignKey:CoachID;references:ID" json:"coach_info"`
}

func (p *EventCoachModel) TableName() string {
	return "event_coach"
}

func (p *EventCoachModelResponse) TableName() string {
	return "event_coach"
}

func (p *EventCoachModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *EventCoachModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
