package main

import "time"

type Participant struct {
	Email string `json:"email"`
	Role string `json:"role"`
	Done bool `json:"done"`
}

type Thing struct {
	Owner Participant `json:"owner"`
	Id int `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Due         time.Time `json:"time"`
	ThingName string `json:"name"`
	ThingLink string `json:"link"`
	IntervalDays int `json:"interval"`
	Participants []Participant `json:"participants"`
}

func (t *Thing) AllDone() bool {
	if len(t.Participants) == 0 {
		return true
	}

	for _,p := range t.Participants {
		if !p.Done {
			return false
		}
	}

	return true
}
