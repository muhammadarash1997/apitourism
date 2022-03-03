package image

type ImageFormatter struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	FileName      string `json:"file_name"`
}

func FormatImageAdded(image Image) ImageFormatter {
	imageFormatted := ImageFormatter{
		ID: image.ID,
		DestinationID: image.DestinationID,
		FileName: image.FileName,
	}

	return imageFormatted
}