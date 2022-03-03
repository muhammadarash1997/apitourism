package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Email        string `gorm:"type:varchar(100)" json:"email"`
	Role         string `gorm:"type:varchar(10)" json:"role"`
	PasswordHash string `gorm:"type:text" json:"password_hash"`
}
