package models

import (
	"time"

	"gorm.io/gorm"
)

type EventPesertaDocumentModel struct {
	ID          NullString `gorm:"column:id;primary_key" json:"id"`
	EventID     NullString `gorm:"column:event_id" json:"event_id"`
	KontingenID NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	PesertaID   NullString `gorm:"column:peserta_id" json:"peserta_id"`
	File        NullString `gorm:"column:file" json:"file"`
	Type        NullString `gorm:"column:type" json:"type"`
	Status      int        `gorm:"column:status;default:0" json:"status"`
	Description NullString `gorm:"column:description" json:"description"`
	INC         int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type EventPesertaDocumentModelResponse = EventPesertaDocumentModel

func (p *EventPesertaDocumentModel) TableName() string {
	return "event_peserta_documents"
}

func (p *EventPesertaDocumentModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventPesertaDocumentModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
