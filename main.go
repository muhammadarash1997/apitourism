package main

import (
	"apitourism/auth"
	"apitourism/bookmark"
	"apitourism/database"
	"apitourism/destination"
	"apitourism/image"
	"apitourism/rating"
	"apitourism/user"
	"apitourism/view"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	userDb := database.StartConnection() // buat objek database

	userRepository := user.NewRepository(userDb)   // buat objek dari repository dengan mempassing objek dari db yg sudah dibuat
	userService := user.NewService(userRepository) // buat objek dari service dengan mempassing objek dari repository yg sudah dibuat
	authService := auth.NewService()
	userHandler := user.NewUserHandler(userService, authService) // buat objek dari userHandler dengan mempassing objek dari service yg sudah dibuat

	destinationRepository := destination.NewRepository(userDb)
	destinationService := destination.NewService(destinationRepository)
	destinationHandler := destination.NewDestinationHandler(destinationService)

	bookmarkRepository := bookmark.NewRepository(userDb)
	bookmarkService := bookmark.NewService(bookmarkRepository)
	bookmarkHandler := bookmark.NewBookmarkHandler(bookmarkService)

	viewRepository := view.NewRepository(userDb)
	viewService := view.NewService(viewRepository)
	viewHandler := view.NewViewHandler(viewService)

	ratingRepository := rating.NewRepository(userDb)
	ratingService := rating.NewService(ratingRepository)
	ratingHandler := rating.NewRatingHandler(ratingService)

	imageRepository := image.NewRepository(userDb)
	imageService := image.NewService(imageRepository)
	imageHandler := image.NewImageHandler(imageService)

	router := gin.Default()

	router.Use(auth.CORSMiddleware())

	// For Admin
	router.POST("/destination", userHandler.AuthenticateHandler, destinationHandler.AddDestinationHandler) // done
	router.POST("/image/:destinationUUID", userHandler.AuthenticateHandler, imageHandler.AddImageHandlerByUUID)
	router.DELETE("/user/:userUUID", userHandler.AuthenticateHandler, userHandler.DeleteUserByUUIDHandler)
	// router.DELETE("/user/:userUUID", userHandler.DeleteUserByUUIDHandler)
	router.DELETE("/destination/:destinationUUID", userHandler.AuthenticateHandler, destinationHandler.DeleteDestinationByUUIDHandler) // done
	router.DELETE("/image/:imageUUID", userHandler.AuthenticateHandler, imageHandler.DeleteImageByUUIDHandler)

	// For User
	router.POST("/user/register", userHandler.RegisterUserHandler)
	router.POST("/user/login", userHandler.LoginHandler)
	router.POST("/bookmark", userHandler.AuthenticateHandler, bookmarkHandler.AddBookmarkHandler)
	router.POST("/view", userHandler.AuthenticateHandler, viewHandler.AddViewHandler)
	router.POST("/rating", userHandler.AuthenticateHandler, ratingHandler.AddRatingHandler)
	router.DELETE("/bookmark/:bookmarkUUID", userHandler.AuthenticateHandler, bookmarkHandler.DeleteBookmarkByUUIDHandler)
	router.GET("/search/destinations/:userCoordinate", userHandler.AuthenticateHandler, destinationHandler.GetAllDestinationsByLimitPageHandler)
	router.GET("/search/destinations/nearby/:userCoordinate", userHandler.AuthenticateHandler, destinationHandler.FindNearbyDestinationHandler)

	router.Run(":8080")
}

// input
// handler mapping input ke struct <-- di sini password masih polos belum di hash dan belum ada createdAt dan updatedAt
// service mapping struct ke struct User <-- di sini password sudah di hash dan ada createdAt dan updatedAt
// repository save struct User ke db
// db

/// CARA INSTANSIASI
// buat objek dari repository dengan mempassing objek dari db yg sudah dibuat
// buat objek dari service dengan mempassing objek dari repository yg sudah dibuat
// buat objek dari userHandler dengan mempassing objek dari service yg sudah dibuat
