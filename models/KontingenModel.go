package models

import (
	"time"

	"gorm.io/gorm"
)

type KontingenModel struct {
	ID     NullString `gorm:"column:id;primary_key" json:"id"`
	UserID NullString `gorm:"column:user_id" json:"user_id"`
	Status int        `gorm:"column:status;default:1" json:"status"`
	INC    int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type KontingenModelResponse struct {
	KontingenModel
	UserInfo UserModelResponse `gorm:"->;foreignKey:UserID;references:ID" json:"user_info"`
}

func (p *KontingenModel) TableName() string {
	return "kontingen"
}

func (p *KontingenModelResponse) TableName() string {
	return "kontingen"
}

func (p *KontingenModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *KontingenModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
