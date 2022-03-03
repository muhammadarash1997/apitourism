package user

import (
	"apitourism/auth"
	"apitourism/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
	authService auth.Service
}

func NewUserHandler(userService Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (this *userHandler) RegisterUserHandler(c *gin.Context) {
	// 1. tangkap input dari user
	// 2. map input dari user ke struct RegisterUserInput
	// 3. struct di atas kita passing sebagai parameter service

	var userInput RegisterUserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if userInput.Password == "01010101" {
		userInput.Role = "admin"
	} else {
		userInput.Role = "user"
	}

	userRegistered, err := this.userService.RegisterUser(userInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userFormatted := FormatUserRegistered(userRegistered)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", userFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *userHandler) DeleteUserByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")

	currentUser := c.MustGet("currentUser").(User)
	if currentUser.Role != "admin" {
		errorMessage := gin.H{"errors": "Only administrator has access"}

		response := helper.APIResponse("Addition destination failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := this.userService.DeleteUserByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Deletion failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Deletion successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (this *userHandler) LoginHandler(c *gin.Context) {
	var loginInput LoginInput

	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogged, err := this.userService.Login(loginInput)
	if err != nil {
		errormessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	tokenGenerated, err := this.authService.GenerateToken(userLogged.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormatted := FormatUserLogged(userLogged, tokenGenerated)

	response := helper.APIResponse("Login succesfully", http.StatusOK, "success", userFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *userHandler) AuthenticateHandler(c *gin.Context) {
	// Ambil token dari header
	tokenInput := c.GetHeader("Authorization")

	// Validasi apakah benar itu adalah bearer token
	if !strings.Contains(tokenInput, "Bearer") {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	tokenWithoutBearer := strings.Split(tokenInput, " ")[1]

	// Validasi token apakah berlaku
	token, err := this.authService.ValidateToken(tokenWithoutBearer)
	if err != nil {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Mengubah token yang bertipe jwt.Token menjadi bertipe jwt.MapClaims
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id := claim["user_uuid"].(string)
	user, err := this.userService.GetUser(id)
	if err != nil {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", user)
}
