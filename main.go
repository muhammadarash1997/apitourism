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
	"net/http"

	_ "apitourism/docs" // This line is necessary for go-swagger to find your docs!

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	var (
		userDb = database.StartConnection()

		userRepository = user.NewRepository(userDb)
		userService    = user.NewService(userRepository)
		authService    = auth.NewService()
		userHandler    = user.NewUserHandler(userService, authService)

		destinationRepository = destination.NewRepository(userDb)
		destinationService    = destination.NewService(destinationRepository)
		destinationHandler    = destination.NewDestinationHandler(destinationService)

		bookmarkRepository = bookmark.NewRepository(userDb)
		bookmarkService    = bookmark.NewService(bookmarkRepository)
		bookmarkHandler    = bookmark.NewBookmarkHandler(bookmarkService)

		viewRepository = view.NewRepository(userDb)
		viewService    = view.NewService(viewRepository)
		viewHandler    = view.NewViewHandler(viewService)

		ratingRepository = rating.NewRepository(userDb)
		ratingService    = rating.NewService(ratingRepository)
		ratingHandler    = rating.NewRatingHandler(ratingService)

		imageRepository = image.NewRepository(userDb)
		imageService    = image.NewService(imageRepository)
		imageHandler    = image.NewImageHandler(imageService)
	)

	router := gin.Default()

	router.Use(cors.Default())

	router.Static("/api/docs", "./swagger")

	// swagger:route GET /api/test test testAPI
	// Test API
	//
	// responses:
	//		200: testAPI200

	// For API Testing
	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// For Admin
	router.POST("/api/destination", userHandler.AuthenticateHandler, destinationHandler.AddDestinationHandler)
	router.POST("/api/image/:destinationUUID", userHandler.AuthenticateHandler, imageHandler.AddImageHandlerByUUID)
	router.DELETE("/api/user/:userUUID", userHandler.AuthenticateHandler, userHandler.DeleteUserByUUIDHandler)
	router.DELETE("/api/destination/:destinationUUID", userHandler.AuthenticateHandler, destinationHandler.DeleteDestinationByUUIDHandler)
	router.DELETE("/api/image/:imageUUID", userHandler.AuthenticateHandler, imageHandler.DeleteImageByUUIDHandler)

	// For User
	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)
	router.POST("/api/bookmark", userHandler.AuthenticateHandler, bookmarkHandler.AddBookmarkHandler)
	router.POST("/api/view", userHandler.AuthenticateHandler, viewHandler.AddViewHandler)
	router.POST("/api/rating", userHandler.AuthenticateHandler, ratingHandler.AddRatingHandler)
	router.DELETE("/api/bookmark/:bookmarkUUID", userHandler.AuthenticateHandler, bookmarkHandler.DeleteBookmarkByUUIDHandler)
	router.GET("/api/destinations/:userCoordinate", userHandler.AuthenticateHandler, destinationHandler.GetAllDestinationsByLimitPageHandler)
	router.GET("/api/destinations/nearby/:userCoordinate", userHandler.AuthenticateHandler, destinationHandler.FindNearbyDestinationHandler)

	router.Run()
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
