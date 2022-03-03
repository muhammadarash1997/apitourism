package bookmark

type BookmarkInput struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	DestinationID string `json:"destination_id" binding:"required,uuid"`
}