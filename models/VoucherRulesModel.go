package models

import (
	"time"

	"gorm.io/gorm"
)

type VoucherRulesModel struct {
	ID                           NullString `gorm:"column:id;primary_key" json:"id"`
	Name                         NullString `gorm:"column:name" json:"name"`
	MaximumRedeemType            int        `gorm:"column:maximum_redeem_type;default:0" json:"maximum_redeem_type"`
	MinimumTotalOrder            int        `gorrm:"column:minimum_total_order;default:0" json:"minimum_total_order"`
	MinimumTotalProduct          int        `gorm:"column:minimum_total_product;default:0" json:"minimum_total_product"`
	MaximumRedeem                int        `gorm:"column:maximum_redeem;default:0" json:"maximum_redeem"`
	UserMinimumTransaction       int        `gorm:"column:user_minimum_transaction;default:0" json:"user_minimum_transaction"`
	MaximumUsagePerUser          int        `gorm:"column:maximum_usage_per_user;default:0" json:"maximum_usage_per_user"`
	MaximumVoucherPerTransaction int        `gorm:"column:maximum_voucher_per_transaction;default:1" json:"maximum_voucher_per_transaction"`
	NewUserOnly                  int        `gorm:"column:new_user_only;default:0" json:"new_user_only"`
	Status                       int        `gorm:"column:status;default:1" json:"status"`
	CreatedFields
	UpdatedFields
}

type VoucherRulesModelResponse = VoucherRulesModel

func (p *VoucherRulesModel) TableName() string {
	return "voucher_rules"
}

func (p *VoucherRulesModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *VoucherRulesModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
