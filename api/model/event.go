package model

type Event struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"eventDate"`
}
