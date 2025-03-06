package models

import "gorm.io/gorm"

type RoleModel struct {
	ID     int        `gorm:"<-:false;column:id;primary_key" json:"id"`
	Role   NullString `gorm:"column:role" json:"role"`
	Module int        `gorm:"column:module;default:0" json:"module"`
}

type RoleModelResponse = RoleModel

func (p *RoleModel) TableName() string {
	return "role"
}

func (p *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *RoleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
