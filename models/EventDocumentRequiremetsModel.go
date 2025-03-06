package models

import (
	"time"

	"gorm.io/gorm"
)

type EventDocumentRequiremetsModel struct {
	ID        NullString `gorm:"column:id;primary_key" json:"id"`
	EventID   NullString `gorm:"column:event_id" json:"event_id"`
	Ijazah    int        `gorm:"column:ijazah;default:0" json:"ijazah"`
	Rapor     int        `gorm:"column:rapor;default:0" json:"rrapor"`
	Mandat    int        `gorm:"column:mandat;default:0" json:"mandat"`
	KK        int        `gorm:"column:kk;default:0" json:"kk"`
	Sehat     int        `gorm:"column:sehat;default:0" json:"sehat"`
	Izin      int        `gorm:"column:izin;default:0" jso:"izin"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy NullString `gorm:"column:updated_by" json:"updated_by"`
}

type EventDocumentRequiremetsModelResponse = EventDocumentRequiremetsModel

func (p *EventDocumentRequiremetsModel) TableName() string {
	return "event_document_requirements"
}

func (p *EventDocumentRequiremetsModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *EventDocumentRequiremetsModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
