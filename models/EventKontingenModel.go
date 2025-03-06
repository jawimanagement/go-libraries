package models

import (
	"time"

	"gorm.io/gorm"
)

type EventKontingenModel struct {
	ID          		NullString	`gorm:"column:id;primary_key" json:"id"`
	EventID				NullString	`gorm:"column:event_id" json:"event_id"`
	Name        		NullString	`gorm:"column:name" json:"name"`
	Status				int			`gorm:"column:status;default:0" json:"status"`
	Paid				int			`gorm:"column:paid;default:0" json:"paid"`
	InvoicePaymentID	NullString	`gorm:"column:invoice_payment_id" json:"invoice_payment_id"`
	INC					int			`gorm:"<-:false;column:inc" json:"inc"`
	CreatedAt   		*time.Time 	`gorm:"column:created_at" json:"created_at"`
	CreatedBy   		NullString 	`gorm:"column:created_by" json:"created_by"`
	UpdatedAt   		*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   		NullString 	`gorm:"column:updated_by" json:"updated_by"`
}

type EventKontingenModelResponse struct {
	ID          		NullString	`gorm:"column:id;primary_key" json:"id"`
	EventID				NullString	`gorm:"column:event_id" json:"event_id"`
	Name        		NullString	`gorm:"column:name" json:"name"`
	Status				int			`gorm:"column:status;default:0" json:"status"`
	Paid				int			`gorm:"column:paid;default:0" json:"paid"`
	InvoicePaymentID	NullString	`gorm:"column:invoice_payment_id" json:"invoice_payment_id"`
	INC					int			`gorm:"<-:false;column:inc" json:"inc"`
	// InvoicePaymentInfo	
	EventInfo			EventsModelResponse	`gorm:"<-;foreignKey:ID;references:EventID" json:"event_info"`
	CreatedAt   		*time.Time 	`gorm:"column:created_at" json:"created_at"`
	CreatedBy   		NullString 	`gorm:"column:created_by" json:"created_by"`
	UpdatedAt   		*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   		NullString 	`gorm:"column:updated_by" json:"updated_by"`
}

func (p *EventKontingenModel) TableName() string {
	return "event_kontingen"
}

func (p *EventKontingenModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *EventKontingenModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}