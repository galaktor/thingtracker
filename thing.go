package main

import "time"

type Thing struct {
	Id string `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Due         time.Time `json:"time"`
	ThingName string `json:"name"`
	ThingLink string `json:"link"`
}
