package image

import (
	"apitourism/helper"
	"apitourism/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageService Service
}

func NewImageHandler(imageService Service) *imageHandler {
	return &imageHandler{imageService}
}

func (this *imageHandler) AddImageHandlerByUUID(c *gin.Context) {
	// Read payload
	// Authorization admin
	// Call process
	// Ouput
	
	// Read payload
	uuid := c.Params.ByName("destinationUUID")

	var imageInput ImageInput
	imageInput.DestinationID = uuid

	// Authorization admin
	currentUser := c.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		errorMessage := gin.H{"errors": "Only administrator has access"}

		response := helper.APIResponse("Addition image failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Read payload file image
	file, err := c.FormFile("image")
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Addition image failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create path storage for store image file and add id to file name
	path := fmt.Sprintf("images storage/%s-%s", imageInput.DestinationID, file.Filename)

	// Save image name in database
	imageAdded, err := this.imageService.AddImage(imageInput, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Addition image failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Store image file to path storage
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Addition image failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	imageFormatted := FormatImageAdded(imageAdded)

	response := helper.APIResponse("Image has been added", http.StatusOK, "success", imageFormatted)
	c.JSON(http.StatusOK, response)
}

func (this *imageHandler) DeleteImageByUUIDHandler(c *gin.Context) {
	// Read payload
	// Authorization admin
	// Call process
	// Output

	uuid := c.Params.ByName("imageUUID")

	currentUser := c.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" {
		errorMessage := gin.H{"errors": "Only administrator has access"}

		response := helper.APIResponse("Deletion image failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := this.imageService.DeleteImageByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		
		response := helper.APIResponse("Deletion Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Deletion successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
