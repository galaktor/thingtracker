package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list", 302)
}

func List(w http.ResponseWriter, r *http.Request) {
	if err := refreshThings(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_list, things)
	case "json": renderJson(w, things)
	}
}

func getThing(r *http.Request) (*Thing,error) {
	vars := mux.Vars(r)
	thingId,_ := strconv.Atoi(vars["thingId"])
	return Get(thingId)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Status", "404")
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "404.html")
}

func View(w http.ResponseWriter, r *http.Request) {
	thing, err := getThing(r)
	if err != nil {
		notFound(w, r)
		return
	}

	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_single, thing)
	case "json": renderJson(w, thing)
	}
}

func EditForm(w http.ResponseWriter, r *http.Request) {
	thing, err := getThing(r)
	if err != nil {
		notFound(w, r)
		return
	}

	renderHtml(w, thing_edit, thing)
}

func EditStore(w http.ResponseWriter, r *http.Request) {
	// TODO syncrhonize in store
	t, err := getThing(r)
	if err != nil {
		notFound(w, r)
		return
	}

	due, err := time.Parse(timeLayout, r.FormValue("due"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Owner = Participant{Email: r.FormValue("owner-email")}
	t.Title = r.FormValue("title")
	t.Description = r.FormValue("description")
	t.Due = due
	t.ThingName = r.FormValue("thingname")
	t.ThingLink = r.FormValue("thinglink")

	parts := []Participant{}
	// stop looping when empty email found
	for i:=0; r.FormValue(fmt.Sprintf("p%v-email", i)) != ""; i++ {
		p := Participant{}
		p.Email = r.FormValue(fmt.Sprintf("p%v-email", i))
		p.Role = r.FormValue(fmt.Sprintf("p%v-role", i))
		p.Done = r.FormValue(fmt.Sprintf("p%v-done", i)) != ""  // no value means unchecked
		parts = append(parts, p)
	}
	t.Participants = parts

	//UPDATE
	if err = Save(t); err != nil {
		http.Error(w, err.Error(), 500)
		return
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
		return
	}
	
	t := &Thing{
		Owner: Participant{Email: r.FormValue("owner-email")},
		Id: strconv.Itoa(getNextId()),
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
		Due: due,
		ThingName: r.FormValue("thingname"),
		ThingLink: r.FormValue("thinglink"),
		Participants: []Participant{},
	}

	p := Participant{}
	// stop looping when empty email found
	for i:=0; r.FormValue(fmt.Sprintf("p%v-email", i)) != ""; i++ {
		p.Email = r.FormValue(fmt.Sprintf("p%v-email", i))
		p.Role = r.FormValue(fmt.Sprintf("p%v-role", i))
		p.Done = r.FormValue(fmt.Sprintf("p%v-done", i)) != ""  // no value means unchecked
		t.Participants = append(t.Participants, p)
	}

	if err = Save(t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	refreshThings()
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


