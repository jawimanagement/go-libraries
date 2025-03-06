package models

import (
	"time"

	"gorm.io/gorm"
)

type EventPlannerModel struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name" json:"name"`
	UserID           NullString `gorm:"column:user_id" json:"user_id"`
	Url              NullString `gorm:"column:url" json:"url"`
	Description      NullString `gorm:"column:description" json:"description"`
	Poster           NullString `gorm:"column:poster" json:"poster"`
	Juknis           NullString `gorm:"column:juknis" json:"juknis"`
	PlannerType      NullInt64  `gorm:"column:planner_type" json:"planner_type"`
	Status           int        `gorm:"column:status;default:1" json:"column:status"`
	BankAccount      NullString `gorm:"column:bank_account" json:"bank_account"`
	BankAccountOwner NullString `gorm:"column:bank_account_owner" json:"bank_account_owner"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
}

type EventPlannerModelResponse struct {
	ID               NullString `gorm:"column:id;primary_key" json:"id"`
	Name             NullString `gorm:"column:name" json:"name"`
	UserID           NullString `gorm:"column:user_id" json:"user_id"`
	Url              NullString `gorm:"column:url" json:"url"`
	Description      NullString `gorm:"column:description" json:"description"`
	Poster           NullString `gorm:"column:poster" json:"poster"`
	Juknis           NullString `gorm:"column:juknis" json:"juknis"`
	PlannerType      NullInt64  `gorm:"column:planner_type" json:"planner_type"`
	Status           int        `gorm:"column:status;default:1" json:"column:status"`
	BankAccount      NullString `gorm:"column:bank_account" json:"bank_account"`
	BankAccountOwner NullString `gorm:"column:bank_account_owner" json:"bank_account_owner"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy        NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *EventPlannerModel) TableName() string {
	return "event_planner"
}

func (p *EventPlannerModel) BeforeCreate(tx *gorm.DB) (err *error) {
	return
}

func (p *EventPlannerModel) BeforeUpdate(tx *gorm.DB) (err *error) {
	return
}
