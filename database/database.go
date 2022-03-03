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
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPass := "0000"
	// dbName := "yourdb"
	
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

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
