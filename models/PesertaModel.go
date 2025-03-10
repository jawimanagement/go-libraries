package models

import (
	"time"

	"gorm.io/gorm"
)

type PesertaModel struct {
	PesertaModelDefault
	CreatedFields
	UpdatedFields
}

type PesertaModelDefault struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name" json:"name"`
	KontingenID      NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	GenderID         NullString `gorm:"column:gender_id" json:"gender_id"`
	Nia              NullString `gorm:"column:nia" json:"nia"`
	Birth            NullDate   `gorm:"column:birth" json:"birth"`
	Domisili         NullString `gorm:"column:domisili" json:"domisili"`
	Nisn             NullString `gorm:"column:nisn" json:"nisn"`
	Photo            NullString `gorm:"column:photo" json:"photo"`
	NIK              NullString `gorm:"column:nik" json:"nik"`
	AktaPhoto        NullString `gorm:"column:akta_photo" json:"akta_photo"`
	INC              int        `gorm:"<-:false;column:inc" json:"inc"`
	Status           int        `gorm:"column:status;default:1" json:"status"`
	PhoneCountryCode NullString `gorm:"column:phone_country_code" json:"phone_country_code"`
	Phone            NullString `gorm:"column:phone" json:"phone"`
}

type PesertaModelResponse struct {
	PesertaModelDefault
	GenderInfo    GenderModelResponse    `gorm:"->;foreignKey:GenderID;references:ID" json:"gender_info"`
	KontingenInfo KontingenModelResponse `gorm:"->;foreignKey:KontingenID;references:ID" json:"kontingen_info"`
}

func (p *PesertaModel) TableName() string {
	return "peserta"
}

func (p *PesertaModelResponse) TableName() string {
	return "peserta"
}

func (p *PesertaModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *PesertaModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
