package bookmark

type BookmarkFormatter struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	DestinationID string `json:"destination_id"`
}

func FormatBookmarkAdded(bookmark Bookmark) BookmarkFormatter {
	bookmarkFormatted := BookmarkFormatter{
		ID: bookmark.ID,
		UserID: bookmark.UserID,
		DestinationID: bookmark.DestinationID,
	}

	return bookmarkFormatted
}