package view

import (
	"gorm.io/gorm"
)

type View struct {
	gorm.Model
	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	DestinationID string `gorm:"type:uuid" json:"destination_id"`
}
