package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"encoding/json"
	"time"

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

func getThing(r *http.Request) *Thing {
	vars := mux.Vars(r)
	thingId,_ := strconv.Atoi(vars["thingId"])
	return things[thingId]
}

func View(w http.ResponseWriter, r *http.Request) {
	thing := getThing(r)
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
	thing := getThing(r)
	if thing == nil {
		http.Error(w, "Thing not found.", 404)
		return
	}

	renderHtml(w, thing_edit, thing)
}

func EditStore(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "POST not implemented", 404)
}

func NewForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	out, err := ioutil.ReadFile("thing_new.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, string(out))
}

var timeLayout = "2006-01-02"
func NewStore(w http.ResponseWriter, r *http.Request) {
	// TODO synchronize, make id fetch and commit atomic
	// consider channel in store, where store calls get next id
	due, err := time.Parse(timeLayout, r.FormValue("due"))
	guard(err)
	
	t := &Thing{
		Id: strconv.Itoa(getNextId()),
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
		Due: due ,
		ThingName: r.FormValue("thingname"),
		ThingLink: r.FormValue("thinglink"),
	}

	filename := fmt.Sprintf("store/%s.thing", t.Id)
	serialized, err := json.Marshal(t)
	guard(err)

	err = ioutil.WriteFile(filename, serialized, os.ModeExclusive)
	fmt.Fprintf(w, "successfully added: %v", filename)
}

func getMimetype(r *http.Request) string {
	switch(r.FormValue("mimetype")) {
	case "json": return "json"
	case "html": return "html"
	default: return "html"
	}
}

