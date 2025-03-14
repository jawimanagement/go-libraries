package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name" json:"name"`
	Email            NullString `gorm:"column:email" json:"email"`
	UserName         NullString `gorm:"column:username" json:"username"`
	Password         NullString `gorm:"column:password,omitempty" json:"password"`
	RoleID           int        `gorm:"column:role_id" json:"role_id"`
	PhoneCountryCode NullString `gorm:"column:phone_country_code" json:"phone_country_code"`
	Phone            NullString `gorm:"column:phone" json:"phone"`
	Status           NullInt64  `gorm:"column:status" json:"status"`
	INC              int        `gorm:"<-:false;column:inc" json:"inc"`
	CreatedFields
	UpdatedFields
}

type UserModelResponse struct {
	UserModel
	RoleInfo RoleModelResponse `gorm:"->;foreignKey:RoleID;references:ID" json:"role_info"`
}

func (p *UserModel) TableName() string {
	return "user"
}

func (p *UserModelResponse) TableName() string {
	return "user"
}

func (p *UserModel) BeforeCreate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *UserModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
