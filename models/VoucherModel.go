package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherModel struct {
	ID           NullString `gorm:"column:id;primary_key" json:"id"`
	ValidFrom    NullDate   `gorm:"column:valid_from" json:"valid_from"`
	ValidUntil   NullDate   `gorm:"column:valid_until" json:"valid_until"`
	VoucherType  int        `gorm:"column:voucher_type;default:1" json:"voucher_type"`
	DiscountType int        `gorm:"column:discount_type;default:1" json:"discount_type"`
	Value        int        `gorm:"column:value" json:"value"`
	Status       int        `gorm:"column:status;default:0" json:"status"`
	CreatedFields
	UpdatedFields
}

type VoucherModelResponse = VoucherModel

func (p *VoucherModel) TableName() string {
	return "voucher"
}

func (p *VoucherModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
