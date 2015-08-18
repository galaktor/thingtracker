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
	http.Redirect(w, r, "/list", 302)
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

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Status", "404")
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "404.html")
}

func View(w http.ResponseWriter, r *http.Request) {
	thing := getThing(r)
	if thing == nil {
		notFound(w, r)
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
		notFound(w, r)
		return
	}

	renderHtml(w, thing_edit, thing)
}

func EditStore(w http.ResponseWriter, r *http.Request) {
	// TODO syncrhonize in store
	t := getThing(r)
	if t == nil {
		notFound(w, r)
		return
	}

	due, err := time.Parse(timeLayout, r.FormValue("due"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	
	t.Title = r.FormValue("title")
	t.Description = r.FormValue("description")
	t.Due = due
	t.ThingName = r.FormValue("thingname")
	t.ThingLink = r.FormValue("thinglink")

	filename := fmt.Sprintf("store/%s.thing", t.Id)
	serialized, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = ioutil.WriteFile(filename, serialized, os.ModeExclusive)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	
	redirUrl := fmt.Sprintf("/show/%v", t.Id)
	http.Redirect(w, r, redirUrl, 302)
}

func NewForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "thing_new.html")
}

var timeLayout = "2006-01-02"
func NewStore(w http.ResponseWriter, r *http.Request) {
	// TODO synchronize, make id fetch and commit atomic
	// consider channel in store, where store calls get next id
	due, err := time.Parse(timeLayout, r.FormValue("due"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	
	t := &Thing{
		Id: strconv.Itoa(getNextId()),
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
		Due: due,
		ThingName: r.FormValue("thingname"),
		ThingLink: r.FormValue("thinglink"),
	}

	filename := fmt.Sprintf("store/%s.thing", t.Id)
	serialized, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = ioutil.WriteFile(filename, serialized, os.ModeExclusive)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	redirUrl := fmt.Sprintf("/show/%v", t.Id)
	http.Redirect(w, r, redirUrl, 302)
}

func getMimetype(r *http.Request) string {
	switch(r.FormValue("mimetype")) {
	case "json": return "json"
	case "html": return "html"
	default: return "html"
	}
}


