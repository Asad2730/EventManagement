package model

import "github.com/google/uuid"

type Event struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventDate   string    `json:"eventDate"`
}
