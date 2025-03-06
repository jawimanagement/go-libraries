package models

import (
	"time"

	"gorm.io/gorm"
)

type EventsModel struct {
	ID                NullString `gorm:"column:id;primary_key" json:"id"`
	Name              NullString `gorm:"column:name" json:"name"`
	UserID            NullString `gorm:"column:user_id" json:"user_id"`
	EventPlannerID    NullString `gorm:"column:event_planner_id" json:"event_planner_id"`
	Url               NullString `gorm:"column:url" json:"url"`
	ShortDesc         NullString `gorm:"column:short_desc" json:"short_desc"`
	DateStart         NullDate   `gorm:"column:date_start" json:"date_start"`
	DateEnd           NullDate   `gorm:"column:date_end" json:"date_end"`
	ProvinceID        NullString `gorm:"column:province_id" json:"province_id"`
	CityID            NullString `gorm:"column:city_id" json:"city_id"`
	DaftarStart       NullDate   `gorm:"column:daftar_start" json:"daftar_start"`
	DaftarEnd         NullDate   `gorm:"column:daftar_end" json:"daftar_end"`
	Logo              NullString `gorm:"column:logo" json:"logo"`
	Poster            NullString `gorm:"column:poster" json:"poster"`
	QRCode            NullString `gorm:"column:qr_code" json:"qr_code"`
	Juknis            NullString `gorm:"column:juknis" json:"juknis"`
	Arena             NullString `gorm:"column:arena" json:"arena"`
	Status            int        `gorm:"column:status;default:0" json:"status"`
	KontingenPrice    float64    `gorm:"column:kontingen_price" json:"kontingen_price"`
	CertificatePrefix NullString `gorm:"column:certificate_prefix" json:"certificate_prefix"`
	Paid              int        `gorm:"column:paid;default:0" json:"paid"`
	IDCard            NullString `gorm:"column:id_card;" json:"id_card"`
	CreatedFields
	UpdatedFields
}

type EventsModelResponse struct {
	EventsModel
	ProvinceInfo     ProvinceModel             `gorm:"->;foreignKey:ProvinceID;references:ProvID" json:"province_info"`
	CityInfo         CityModel                 `gorm:"->;foreignKey:CityID;references:CityID" json:"city_info"`
	UserInfo         UserModelResponse         `gorm:"->;foreignKey:UserID;references:ID" json:"user_info"`
	EventPlannerInfo EventPlannerModelResponse `gorm:"->;foreignKey:EventPlannerID;references:ID" json:"event_planner_info"`
}

func (p *EventsModel) TableName() string {
	return "events"
}

func (p *EventsModelResponse) TableName() string {
	return "events"
}

func (p *EventsModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedAt = &now
	p.UpdatedAt = &now
	return
}

func (p *EventsModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.UpdatedAt = &now
	return
}
