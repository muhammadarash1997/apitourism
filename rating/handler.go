package rating

import (
	"apitourism/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ratingHandler struct {
	ratingService Service
}

func NewRatingHandler(ratingService Service) *ratingHandler {
	return &ratingHandler{ratingService}
}

func (this * ratingHandler) AddRatingHandler(c *gin.Context) {
	// Read payload
	// Call process
	// Output

	var ratingInput RatingInput

	err := c.ShouldBindJSON(&ratingInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Addition rating failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ratingAdded, err := this.ratingService.AddRating(ratingInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Addition rating failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ratingFormatted := FormatRatingAdded(ratingAdded)

	response := helper.APIResponse("Rating has been added", http.StatusOK, "success", ratingFormatted)

	c.JSON(http.StatusOK, response)
}