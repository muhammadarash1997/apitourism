package rating

type RatingInput struct {
	DestinationID string `json:"destination_id" binding:"required,uuid"`	// required artinya data harus diisi, uuid artinya data harus berupa uuid
	UserID string `json:"user_id" binding:"required,uuid"`
	Rate int `json:"rate" binding:"required"`
}