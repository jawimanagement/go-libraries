package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceDataModel struct {
	ID            NullString `gorm:"column:id;primary_key" json:"id"`
	InvoiceID     NullString `gorm:"column:invoice_id" json:"invoice_id"`
	EventID       NullString `gorm:"column:event_id" json:"event_id"`
	PlanID        NullString `gorm:"column:plan_id" json:"plan_id"`
	EventDetailID NullString `gorm:"column:event_detail_id" json:"event_detail_id"`
	AdditionalID  NullString `gorm:"column:additional_id" json:"additional_id"`
	Price         float64    `gorm:"column:price" json:"price"`
	Qty           NullInt64  `gorm:"column:qty" json:"qty"`
	PromoID       NullString `gorm:"column:promo_id" json:"promo_id"`
	PromoInfo     NullString `gorm:"column:promo_info" json:"promo_info"`
	PromoValue    float64    `gorm:"column:promo_value" json:"promo_value"`
	UserInfo      NullString `gorm:"column:user_info" json:"user_info"`
	ProductInfo   NullString `gorm:"column:product_info" json:"product_info"`
	DiscountPrice float64    `gorm:"column:discount_price" json:"discount_price"`
	FinalPrrice   float64    `gorm:"column:final_price" json:"final_price"`
	ProductData   NullString `gorm:"column:product_data" json:"product_data"`
	CreatedFields
	UpdatedFields
}

type InvoiceDataModelResponse = InvoiceDataModel

func (p *InvoiceDataModel) TableName() string {
	return "invoice_data"
}

func (p *InvoiceDataModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *InvoiceDataModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
