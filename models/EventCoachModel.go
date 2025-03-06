package models

import (
	"time"

	"gorm.io/gorm"
)

type EventCoachModel struct {
	ID      NullString `gorm:"column:id;primary_key" json:"id"`
	EventID NullString `gorm:"column:event_id" json:"event_id"`
	CoachID NullString `gorm:"column:coach_id" json:"coach_id"`
	QRCode  NullString `gorm:"column:qrcode" json:"qrcode"`
	INC     int        `gorm:"<-:false;column:inc" json:"inc"`
	Status  int        `gorm:"column:status;default:1" json:"status"`
	CreatedFields
	UpdatedFields
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
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventCoachModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
