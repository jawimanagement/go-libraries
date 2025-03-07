package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherUserModel struct {
	ID        NullString `gorm:"column:id;primary_key" json:"id"`
	VoucherID NullString `gorm:"column:voucher_id" json:"voucher_id"`
	UserID    NullString `gorm:"column:user_id" json:"user_id"`
	UserInfo  NullString `gorm:"column:user_info" json:"user_info"`
	CreatedFields
	UpdatedFields
}

type VoucherUserModelResponse struct {
	ID        NullString `gorm:"column:id;primary_key" json:"id"`
	VoucherID NullString `gorm:"column:voucher_id" json:"voucher_id"`
	UserID    NullString `gorm:"column:user_id" json:"user_id"`
	CreatedFields
	UpdatedFields
	VoucherInfo VoucherModelResponse `gorm:"->;foreignKey:VoucherID;references:ID" json:"voucher_info"`
}

func (p *VoucherUserModel) TableName() string {
	return "voucher_user"
}

func (p *VoucherUserModelResponse) TableName() string {
	return "voucher_user"
}

func (p *VoucherUserModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherUserModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
