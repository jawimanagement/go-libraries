package models

import (
	"time"

	"gorm.io/gorm"
)

type CartModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	UserID           NullString `gorm:"column:user_id" json:"user_id"`
	EventID          NullString `gorm:"column:event_id" json:"event_id"`
	EventKontingenID NullString `gorm:"column:event_kontingen_id" json:"event_kontingen_id"`
	PlanID           NullString `gorm:"column:plan_id" json:"plan_id"`
	EventDetailID    NullString `gorm:"column:event_detail_id" json:"event_detail_id"`
	EventPesertaID   NullString `gorm:"column:event_peserta_id" json:"event_peserta_id"`
	AdditionalID     NullString `gorm:"column:additional_id" json:"additioal_id"`
	Price            float64    `gorm:"column:price;default:0" json:"price"`
	Qty              NullInt64  `gorm:"column:qty" json:"qty"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
}

type CartModelResponse struct {
	CartModel
	EventInfo          EventsModelResponse         `gorm:"->;foreignKey:EventID;references:ID" json:"event_info"`
	EventKontingenInfo EventKontingenModelResponse `gorm:"->;foreignKey:EventKontingenID;references:ID" json:"event_kontingen_info"`
	EventDetailInfo    EventDetailModelResponse    `gorm:"->;foreignKey:EventDetailID;references:ID" json:"event_detail_info"`
	EventPesertaInfo   EventPesertaModelResponse   `gorm:"->;foreignKey:EventPesertaID;references:ID" json:"event_peserta_info"`
}

func (p *CartModel) TableName() string {
	return "cart"
}

func (p *CartModelResponse) TableName() string {
	return "cart"
}

func (p *CartModel) BeforeCreate(tx *gorm.DB) (err *error) {
	return
}

func (p *CartModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	return
}
