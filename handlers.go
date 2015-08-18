package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func List(w http.ResponseWriter, r *http.Request) {
	if err := refreshThings(); err != nil {
		http.Error(w, err.Error(), 500)
	}
	
	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_list, things)
	case "json": renderJson(w, things)
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	thingId,_ := strconv.Atoi(vars["thingId"])
	thing := things[thingId]

	if thing == nil {
		http.Error(w, "Thing not found.", 404)
		return
	}

	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_single, thing)
	case "json": renderJson(w, thing)
	}
}

func EditForm(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "GET not implemented", 404)
}

func EditStore(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "PUT not implemented", 404)
}

func NewForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	out, err := ioutil.ReadFile("thing_new.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, string(out))
}

func NewStore(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "POST not implemented", 404)
}

func getMimetype(r *http.Request) string {
	switch(r.FormValue("mimetype")) {
	case "json": return "json"
	case "html": return "html"
	default: return "html"
	}
}
