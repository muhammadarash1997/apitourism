package image

type ImageInput struct {
	DestinationID string `json:"destination_id" binding:"required,uuid"`
}