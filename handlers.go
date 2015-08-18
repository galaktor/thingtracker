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
	// TODO syncrhonize in store
	t := getThing(r)
	if t == nil {
		http.Error(w, "Thing not found.", 404)
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
	
	//	fmt.Fprintf(w, "successfully updated: %v", t.Id)
	redirUrl := fmt.Sprintf("/show/%v", t.Id)
	http.Redirect(w, r, redirUrl, 302)
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

	//fmt.Fprintf(w, "successfully added: %v", filename)
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

