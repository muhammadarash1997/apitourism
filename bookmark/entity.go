package bookmark

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID        string `gorm:"type:uuid" json:"user_id"`
	DestinationID string `gorm:"type:uuid" json:"destination_id"`
}
