package main

import "time"

type Thing struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Tracker struct {
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Due         time.Time `json:"time"`
	Thing       Thing
}

//type Things []Thing










