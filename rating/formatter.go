package rating

type RatingFormatter struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	UserID        string `json:"user_id"`
	Rate          int    `json:"rate"`
}

func FormatRatingAdded(rating Rating) RatingFormatter {
	ratingFormatted := RatingFormatter{
		ID:            rating.ID,
		DestinationID: rating.DestinationID,
		UserID:        rating.UserID,
		Rate:          rating.Rate,
	}

	return ratingFormatted
}
