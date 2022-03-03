package view

type viewFormatter struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	UserID        string `json:"user_id"`
}

func FormatViewAdded(view View) viewFormatter {
	viewFormatted := viewFormatter{
		ID: view.ID,
		DestinationID: view.DestinationID,
	}

	return viewFormatted
}