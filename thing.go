package main

import "time"

type Participant struct {
	Email string `json:"email"`
	Role string `json:"role"`
	Done bool `json:"done"`
}

type Thing struct {
	Owner Participant `json:"owner"`
	Id string `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Due         time.Time `json:"time"`
	ThingName string `json:"name"`
	ThingLink string `json:"link"`
	Participants []Participant `json:"participants"`
}
