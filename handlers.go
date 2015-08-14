package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Things(w http.ResponseWriter, r *http.Request) {
	things := []Thing{
		Thing{Name: "Foo", Link: "http://www.foo.bar"},
	}

	if err := json.NewEncoder(w).Encode(things); err != nil {
		panic(err)
	}
}

func SingleThing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	thingId := vars["thingId"]
	fmt.Fprintln(w, "Thing id:", thingId)
}










