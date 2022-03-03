package destination

import (
	"apitourism/helper"
	"apitourism/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type destinationHandler struct {
	destinationService Service
}

func NewDestinationHandler(destinationService Service) *destinationHandler {
	return &destinationHandler{destinationService}
}

func (this *destinationHandler) AddDestinationHandler(c *gin.Context) {
	// 1. Read payload
	// 2. Authorization admin
	// 3. Call process function
	// 4. Output
	var destinationInput DestinationInput

	err := c.ShouldBindJSON(&destinationInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Addition destination failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		errorMessage := gin.H{"errors": "Only administrator has access"}

		response := helper.APIResponse("Addition destination failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	destinationAdded, err := this.destinationService.AddDestination(destinationInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Addition destination failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	destinationFormatted := FormatDestinationAdded(destinationAdded)

	response := helper.APIResponse("Destination has been added", http.StatusOK, "success", destinationFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *destinationHandler) DeleteDestinationByUUIDHandler(c *gin.Context) {
	// Read payload
	// Authorization admin
	// Call process
	// Output

	uuid := c.Params.ByName("destinationUUID")

	currentUser := c.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		errorMessage := gin.H{"errors": "Only administrator has access"}

		response := helper.APIResponse("Deletion destination failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := this.destinationService.DeleteDestinationByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Deletion failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Deletion successfully", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)
}

func (this *destinationHandler) GetAllDestinationsByLimitPageHandler(c *gin.Context) {
	// Get path params
	// Get query params
	// Call process
	// Output

	// Get path params
	userCoordinate := c.Params.ByName("userCoordinate")

	// Get query params
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}

		response := helper.APIResponse("Get destinations failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}

		response := helper.APIResponse("Get destinations failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call process
	destinations, err := this.destinationService.GetAllDestinationsByLimitPage(limit, page)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}

		response := helper.APIResponse("Get destinations failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	destinationsFormatted := FormatDestinationsGotten(userCoordinate, destinations)

	response := helper.APIResponse("Get destinations successfully", http.StatusOK, "success", destinationsFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *destinationHandler) FindNearbyDestinationHandler(c *gin.Context) {
	// Read path param & query param
	// Call process
	// Output

	// Read path param & query param
	userCoordinate := c.Params.ByName("userCoordinate")
	inputType := c.Query("type")
	inputRating, _ := strconv.ParseBool(c.Query("rating"))
	inputPopularity, _ := strconv.ParseBool(c.Query("popularity"))

	// Call process
	nearbyDestinationsGotten, err := this.destinationService.FindNearbyDestination(userCoordinate, inputType, inputRating, inputPopularity)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get destinations failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	nearbyDestinationsFormatted := FormatNearbyDestinationsGotten(userCoordinate, nearbyDestinationsGotten)

	response := helper.APIResponse("Get destinations successfully", http.StatusOK, "success", nearbyDestinationsFormatted)
	c.JSON(http.StatusOK, response)
}
