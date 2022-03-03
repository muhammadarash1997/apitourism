package view

import (
	"apitourism/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type viewHandler struct {
	viewService Service
}

func NewViewHandler(service Service) *viewHandler {
	return &viewHandler{service}
}

func (this *viewHandler) AddViewHandler(c *gin.Context) {
	// Read payload
	// Call process
	// Output

	var viewInput ViewInput

	err := c.ShouldBindJSON(&viewInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Addition view failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	viewAdded, err := this.viewService.AddView(viewInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Addition view failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
	}

	viewFormatted := FormatViewAdded(viewAdded)
	response := helper.APIResponse("View has been added", http.StatusOK, "success", viewFormatted)

	c.JSON(http.StatusOK, response)
}
