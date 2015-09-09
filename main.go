package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO PUT IN CONFIG
const (
	URL_IP   = "163.33.212.69"
	URL_PORT = "8080"
	URL_ROOT = "http://163.33.212.69:8080"
)

func main() {
	router := NewRouter()
	startTimer()
	log.Fatal(http.ListenAndServe(fmt.Sprint(":",URL_PORT), router))
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}
