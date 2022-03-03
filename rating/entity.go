package rating

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	DestinationID string `gorm:"type:uuid" json:"destination_id"`
	UserID        string `gorm:"type:uuid" json:"user_id"`
	Rate          int    `gorm:"type:int" json:"rate"`
}
