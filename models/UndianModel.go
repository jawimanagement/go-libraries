package models

import (
	"time"

	"gorm.io/gorm"
)

type UndianModel struct {
	ID              NullString `gorm:"column:id;primary_key" json:"id"`
	EventID         NullString `gorm:"column:event_id" json:"event_id"`
	KategoriID      NullString `gorm:"column:kategori_id" json:"kategori_id"`
	GolonganID      NullString `gorm:"column:golongan_id" json:"golongan_id"`
	GenderID        NullString `gorm:"column:gender_id" json:"gender_id"`
	KelasTandingID  NullString `gorm:"kelas_id" json:"kelas_id"`
	KontingenID     NullString `gorm:"column:kontingen_id" json:"kontingen_id"`
	GroupID         NullString `gorm:"column:group_id" json:"group_id"`
	NoUndi          int64      `gorm:"column:noundi" json:"noundi"`
	Corner          int        `gorm:"column:corner" json:"corner"`
	PartaiReference NullString `gorm:"partai_reference" json:"partai_reference"`
	ScheduleLevel   int        `gorm:"column:schedule_level" json:"schedule_level"`
	UndianType      int        `gorm:"column:undian_type;default:1" json:"undian_type"`
	Pool            int        `gorm:"column:pool;default:1" json:"pool"`
	UndianID        NullString `gorm:"column:undian_id" json:"undian_id"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy       NullString `gorm:"column:created_by" json:"created_by"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy       NullString `gorm:"column:updated_by" json:"updated_by"`
}

func (p *UndianModel) TableName() string {
	return "undian"
}

func (p *UndianModel) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (p *UndianModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (p *UndianModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (p *UndianModel) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
