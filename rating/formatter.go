package rating

type ratingFormatter struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	UserID        string `json:"user_id"`
	Rate          int    `json:"rate"`
}

func FormatRatingAdded(rating Rating) ratingFormatter {
	ratingFormatted := ratingFormatter{
		ID:            rating.ID,
		DestinationID: rating.DestinationID,
		UserID:        rating.UserID,
		Rate:          rating.Rate,
	}

	return ratingFormatted
}
