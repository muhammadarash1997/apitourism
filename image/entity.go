package image

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	DestinationID string `gorm:"type:uuid" json:"destination_id"`
	FileName      string `gorm:"type:varchar(100)" json:"file_name"`
}
