package models

import (
	"time"

	"gorm.io/gorm"
)

type ManagerModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name" json:"name"`
	KontingenID      NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	GenderID         NullString `gorm:"column:gender_id" json:"gender_id"`
	PhoneCountryCode NullString `gorm:"column:phone_country_code" json:"phone_country_code"`
	Phone            NullString `gorm:"column:phone" json:"phone"`
	Photo            NullString `gorm:"column:photo" json:"photo"`
	CreatedFields
	UpdatedFields
}

type ManagerModelResponse struct {
	ManagerModel
	KontingenInfo KontingenModelResponse `gorm:"->;foreignKey:KontingenID;references:ID" json:"kontingen_info"`
	GenderInfo    GenderModelResponse    `gorm:"->;foreignKey:GenderID;references:ID" json:"gender_info"`
}

func (p *ManagerModel) TableName() string {
	return "manager"
}

func (p *ManagerModelResponse) TableName() string {
	return "manager"
}

func (p *ManagerModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *ManagerModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
