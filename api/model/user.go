package model

type User struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	EventId string `json:"eventId"`
}
