package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherUsageModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	UserID      NullString `gorm:"column:user_id" json:"user_id"`
	UserInfo    NullString `gorm:"column:user_info" json:"user_info"`
	VoucherID   NullString `gorm:"column:voucher_id" json:"voucher_id"`
	InvoiceID   NullString `gorm:"column:invoice_id" json:"invoice_id"`
	VoucherInfo NullString `gorm:"column:voucher_info" json:"voucher_info"`
	Status      int        `gorm:"column:status;default:0" json:"status"`
	CreatedFields
	UpdatedFields
}

type VoucherUsageModelResponse struct {
	ID          NullString           `gorm:"column:id;primary_key" json:"id"`
	UserID      NullString           `gorm:"column:user_id" json:"user_id"`
	UserInfo    UserModelResponse    `gorm:"->;foreignKey:UserID;references:ID" json:"user_info"`
	VoucherID   NullString           `gorm:"column:voucher_id" json:"voucher_id"`
	InvoiceID   NullString           `gorm:"column:invoice_id" json:"invoice_id"`
	VoucherInfo VoucherModelResponse `gorm:"->;foreignKey:VoucherID;references:ID" json:"voucher_info"`
	Status      int                  `gorm:"column:status;default:0" json:"status"`
	CreatedFields
	UpdatedFields
}

func (p *VoucherUsageModel) TableName() string {
	return "voucher_usage"
}

func (p *VoucherUsageModelResponse) TableName() string {
	return "voucher_usage"
}

func (p *VoucherUsageModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherUsageModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
