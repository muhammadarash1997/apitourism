package destination

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Type    string `gorm:"type:varchar(100)" json:"type"`
	GeoHash string `gorm:"type:varchar(100)" json:"geo_hash"`
}

type NearbyDestination struct {
	ID string
	Name string
	Type string
	GeoHash string
	ViewAmount int
	Rating float64
}