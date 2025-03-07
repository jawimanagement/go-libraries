package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherDataModel struct {
	ID             NullString `gorm:"column:id;primary_key" json:"id"`
	VoucherID      NullString `gorm:"column:voucher_id" json:"voucher_id"`
	VoucherRulesID NullString `gorm:"column:voucher_rules_id" json:"voucher_rules_id"`
	CreatedFields
	UpdatedFields
}

type VoucherDataModelResponse struct {
	VoucherDataModel
	VoucherInfo      VoucherModelResponse      `gorm:"->;foreignKey:VoucherID;references:ID" json:"voucher_info"`
	VoucherRulesInfo VoucherRulesModelResponse `gorm:"->;foreignKey:VoucherRulesID;references:ID" json:"voucher_rules_info"`
}

func (p *VoucherDataModel) TableName() string {
	return "voucher_data"
}

func (p *VoucherDataModelResponse) TableName() string {
	return "voucher_data"
}

func (p *VoucherDataModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherDataModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
