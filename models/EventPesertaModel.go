package models

type EventPesertaModel struct {
	ID          		NullString	`gorm:"column:id;primary_key" json:"id"`
	EventID				NullString	`gorm:"column:event_id" json:"event_id"`
	KontingenID			NullString	`gorm:"column:kontingen_id" json:"kontingen_id"`
	EventDetailID		NullString	`gorm:"column:event_detail_id" json:"event_detail_id"`
	
}