package models

import (
	"time"

	"gorm.io/gorm"
)

type CoachModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	KontingenID      NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	GenderID         NullString `gorm:"column:gender_id" json:"gender_id"`
	Nia              NullString `gorm:"column:nia" json:"nia"`
	Name             NullString `gorm:"column:name" json:"name"`
	PhoneCountryCode NullString `gorm:"column:phone_country_code" json:"phone_country_code"`
	Phone            NullString `gorm:"column:phone" json:"phone"`
	Photo            NullString `gorm:"column:photo" json:"photo"`
	Status           int        `gorm:"column:status;default:1" json:"status"`
	INC              int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type CoachModelResponse struct {
	CoachModel
	KontingenInfo KontingenModelResponse `gorm:"->;foreignKey:KontingenID;references:ID" json:"kontingen_info"`
	GenderInfo    GenderModelResponse    `gorm:"->;foreignKey:GenderID;references:ID" json:"gender_info"`
}

func (p *CoachModel) TableName() string {
	return "coach"
}

func (p *CoachModelResponse) TableName() string {
	return "coach"
}

func (p *CoachModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *CoachModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
