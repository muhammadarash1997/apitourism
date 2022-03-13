package bookmark

import (
	"apitourism/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookmarkHandler struct {
	bookmarkService Service
}

func NewBookmarkHandler(bookmarkService Service) *bookmarkHandler {
	return &bookmarkHandler{bookmarkService}
}

// swagger:route POST /api/bookmark bookmark addBookmark
// Add a bookmark of destination. This can only be done by the logged in user
//
// Security:
// - Bearer:
// responses:
//		200: bookmarkAdded200
//		422: errorResponse

func (this *bookmarkHandler) AddBookmarkHandler(c *gin.Context) {
	// 1. Read payload
	// 2. Call process function
	// 3. Output

	var bookmarkInput BookmarkInput

	err := c.ShouldBindJSON(&bookmarkInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Addition bookmark failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bookmarkAdded, err := this.bookmarkService.AddBookmark(bookmarkInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Addition destination failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bookmarkFormatted := FormatBookmarkAdded(bookmarkAdded)

	response := helper.APIResponse("Bookmark has been added", http.StatusOK, "success", bookmarkFormatted)

	c.JSON(http.StatusOK, response)
}

// swagger:route DELETE /api/bookmark/{bookmarkUUID} bookmark deleteBookmark
// Delete a bookmark of destination. This can only be done by the logged in user
//
// Security:
// - Bearer:
// responses:
//		200: bookmarkDeleted200
//		400: errorResponse

func (this *bookmarkHandler) DeleteBookmarkByUUIDHandler(c *gin.Context) {
	// Read payload
	// Call process
	// Output

	uuid := c.Params.ByName("bookmarkUUID")

	err := this.bookmarkService.DeleteBookmarkByUUID(uuid)
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Deletion failed", http.StatusBadRequest, "error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Deletion success", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)
}
