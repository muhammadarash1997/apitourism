package bookmark

type bookmarkFormatter struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	DestinationID string `json:"destination_id"`
}

func FormatBookmarkAdded(bookmark Bookmark) bookmarkFormatter {
	bookmarkFormatted := bookmarkFormatter{
		ID: bookmark.ID,
		UserID: bookmark.UserID,
		DestinationID: bookmark.DestinationID,
	}

	return bookmarkFormatted
}