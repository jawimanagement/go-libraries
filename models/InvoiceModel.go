package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceModel struct {
	ID            NullString `gorm:"column:id;primary_key" json:"id"`
	ShortCode     NullString `gorm:"column:short_core" json:"short_code"`
	UserID        NullString `gorm:"column:user_id" json:"user_id"`
	GrandTotal    float64    `gorm:"column:grand_total" json:"grand_total"`
	SenderInfo    NullString `gorm:"column:sender_info" json:"sender_info"`
	Status        int        `gorm:"column:status;default:1" json:"status"`
	INC           int        `gorm:"<-:false;column:inc" json:"inc"`
	TotalDiscount float64    `gorm:"column:total_discount" json:"total_discount"`
	AdminFee      float64    `gorm:"column:admin_fee" json:"admin_fee"`
	SumGrandTotal float64    `gorm:"column:sum_grand_total" json:"sum_grand_total"`
	CreatedFields
	UpdatedFields
}

type InvoiceModelResponse struct {
	InvoiceModel
	UserInfo UserModelResponse `gorm:"->;foreignKey:UserID;references:ID" json:"user_info"`
}

func (p *InvoiceModel) TableName() string {
	return "invoice"
}

func (p *InvoiceModelResponse) TableName() string {
	return "invoice"
}

func (p *InvoiceModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *InvoiceModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
