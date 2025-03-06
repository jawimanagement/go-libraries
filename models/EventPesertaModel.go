package models

import (
	"time"

	"gorm.io/gorm"
)

type EventPesertaModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	EventID          NullString `gorm:"column:event_id" json:"event_id"`
	KontingenID      NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	EventDetailID    NullString `gorm:"column:event_detail_id" json:"event_detail_id"`
	PesertaID        NullString `gorm:"column:peserta_id" json:"peserta_id"`
	QRCode           NullString `gorm:"column:qrcode" json:"qrcode"`
	INC              int        `gorm:"<-:false;column:inc" json:"inc"`
	GroupID          NullString `gorm:"column:group_id" json:"group_id"`
	Status           int        `gorm:"column:status;default:0" json:"status"`
	Paid             int        `gorm:"column:paid;default:0" json:"paid"`
	InvoicePaymentID NullString `gorm:"column:invoice_payment_id" json:"invoice_payment_id"`
	CreatedFields
	UpdatedFields
}

type EventPesertaModelResponse struct {
	EventPesertaModel
	InvoicePaymentInfo InvoicePaymentModelResponse `gorm:"->;foreignKey:InvoicePaymentID;references:ID" json:"invoice_payment_info"`
	PesertaInfo        PesertaModelResponse        `gorm:"->;foreignKey:PesertaID;references:ID" json:"peserta_info"`
	EventDetailInfo    EventDetailModelResponse    `gorm:"->;foreignKey:EventDetailID;references:ID" json:"event_detail_info"`
	KontingenInfo      EventKontingenModelResponse `gorm:"->;foreignKey:KontingenID;references:ID" json:"kontingen_info"`
	EventInfo          EventsModelResponse         `gorm:"->;foreignKey:EventID;references:ID" json:"event_info"`
}

func (p *EventPesertaModel) TableName() string {
	return "event_peserta"
}

func (p *EventPesertaModelResponse) TableName() string {
	return "event_peserta"
}

func (p *EventPesertaModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventPesertaModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
