package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoicePaymentModel struct {
	ID            NullString `gorm:"column:id;primary_key" json:"id"`
	InvoiceID     NullString `gorm:"column:invoice_id" json:"invoice_id"`
	PaymentInfo   NullString `gorm:"column:payment_info" json:"payment_info"`
	PayerInfo     NullString `gorm:"column:payer_info" json:"payer_info"`
	Value         float64    `gorm:"column:value" json:"value"`
	PaidOn        *time.Time `gorm:"column:paid_on" json:"paid_on"`
	MaxTimePay    *time.Time `gorm:"column:max_time_pay" json:"max_time_pay"`
	PaymentStatus int        `gorm:"column:payment_status;default:0" json:"payment_status"`
	Notes         NullString `gorm:"column:notes" json:"notes"`
	ExternalID    NullString `gorm:"column:external_id" json:"external_id"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy     NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy     NullString `gorm:"column:updated_by" json:"updated_by"`
}

type InvoicePaymentModelResponse struct {
	InvoicePaymentModel
	InvoiceInfo InvoiceModelResponse `gorm:"->;foreignKey:InvoiceID;references:ID" json:"invoice_info"`
}

func (p *InvoicePaymentModel) TableName() string {
	return "invoice_payment"
}

func (p *InvoicePaymentModelResponse) TableName() string {
	return "invoice_payment"
}

func (p *InvoicePaymentModel) BeforeCreate(tx *gorm.DB) (err *error) {
	return
}

func (p *InvoicePaymentModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	return
}
