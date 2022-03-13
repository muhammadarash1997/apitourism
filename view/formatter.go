package view

type ViewFormatter struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	UserID        string `json:"user_id"`
}

func FormatViewAdded(view View) ViewFormatter {
	viewFormatted := ViewFormatter{
		ID: view.ID,
		DestinationID: view.DestinationID,
	}

	return viewFormatted
}