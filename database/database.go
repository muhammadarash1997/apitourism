package database

import (
	"fmt"
	"log"
	"os"

	"apitourism/bookmark"

	"apitourism/destination"
	"apitourism/image"
	"apitourism/rating"
	"apitourism/user"
	"apitourism/view"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// jika menggunakan heroku maka sslmode harus require (sslmode=require), jika tidak maka sslmode=disable
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to database")
		return nil
	}
	fmt.Println("Success to connect to database")

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&destination.Destination{})
	db.AutoMigrate(&view.View{})
	db.AutoMigrate(&rating.Rating{})
	db.AutoMigrate(&image.Image{})
	db.AutoMigrate(&bookmark.Bookmark{})
	return db
}
