package models

import (
	"time"

	"gorm.io/gorm"
)

type EventKontingenModel struct {
	EventKontingenModelDefault
	CreatedFields
	UpdatedFields
}

type EventKontingenModelDefault struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	EventID          NullString `gorm:"column:event_id" json:"event_id"`
	Name             NullString `gorm:"column:name" json:"name"`
	Status           int        `gorm:"column:status;default:0" json:"status"`
	Paid             int        `gorm:"column:paid;default:0" json:"paid"`
	InvoicePaymentID NullString `gorm:"column:invoice_payment_id" json:"invoice_payment_id"`
	INC              int        `gorm:"<-:false;column:inc" json:"inc"`
}

type EventKontingenModelResponse struct {
	EventKontingenModelDefault
	EventInfo EventsModelResponse `gorm:"->;foreignKey:EventID;references:ID" json:"event_info"`
}

func (p *EventKontingenModel) TableName() string {
	return "event_kontingen"
}

func (p *EventKontingenModelResponse) TableName() string {
	return "event_kontingen"
}

func (p *EventKontingenModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventKontingenModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
