package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	startTimer()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}
