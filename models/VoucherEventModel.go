package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherEventModel struct {
	ID        NullString `gorm:"column:id;primary_key" json:"id"`
	VoucherID NullString `gorm:"column:voucher_id" json:"voucher_id"`
	EventID   NullString `gorm:"column:event_id" json:"event_id"`
	EventInfo NullString `gorm:"column:event_info" json:"event_info"`
	CreatedFields
	UpdatedFields
}

type VoucherEventModelResponse struct {
	VoucherEventModel
	VoucherInfo VoucherModelResponse `gorm:"->;foreignKey:VoucherID;references:ID" json:"voucher_info"`
}

func (p *VoucherEventModel) TableName() string {
	return "voucher_event"
}

func (p *VoucherEventModelResponse) TableName() string {
	return "voucher_event"
}

func (p *VoucherEventModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherEventModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
