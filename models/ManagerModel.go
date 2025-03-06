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
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
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
	return
}

func (p *ManagerModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
