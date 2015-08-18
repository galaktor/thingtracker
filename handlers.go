package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"io/ioutil"
	"path/filepath"

	"github.com/gorilla/mux"
)

var thing_list *template.Template
var thing_single *template.Template

var things []Thing

func init() {
	loadTemplates()
	t, err := deserThings()
	guard(err)
	things = t
	fmt.Println(things)
}

func loadTemplates() {
	file, err := ioutil.ReadFile("thing_single.tmpl")
	guard(err)
	thing_single, err = template.New("thing_single").Parse(string(file))
	guard(err)

	file, err = ioutil.ReadFile("thing_list.tmpl")
	guard(err)
	thing_list, err = template.New("thing_list").Parse(string(file))
	guard(err)
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}

func deserThings() ([]Thing, error) {
	files, err := ioutil.ReadDir("store")
	guard(err)

	out := []Thing{}
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".thing" {
				fullpath := "store/" + file.Name()
				content, err := ioutil.ReadFile(fullpath)
				guard(err)
				t := Thing{}
				err = json.Unmarshal(content, &t)
				guard(err)
				out = append(out, t)
			}
		}
	}
	
	return out, err	
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Reload(w http.ResponseWriter, r *http.Request) {
	var err error
	if things, err = deserThings(); err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintln(w, "Reload complete.")
}

func Things(w http.ResponseWriter, r *http.Request) {
	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_list, things)
	case "json": renderJson(w, things)
	}
}

func renderJson(out http.ResponseWriter, i interface{}) {
	out.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(out).Encode(i); err != nil {
		http.Error(out, err.Error(), 500)
	}
}

func renderHtml(out http.ResponseWriter, t *template.Template, i interface{}) {
	out.Header().Set("Content-Type", "text/html")
	if err := t.Execute(out, i); err != nil {
		http.Error(out, err.Error(), 500)
	}
}

func getMimetype(r *http.Request) string {
	switch(r.FormValue("mimetype")) {
	case "json": return "json"
	case "html": return "html"
	default: return "html"
	}
}

func SingleThing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	thingId := vars["thingId"]
	thing := &Thing{Id: thingId,
		Title:       "foo",
		Description: "bar",
		Due:         time.Now(),
		ThingName:   "foo name",
		ThingLink:   "http://foo.bar"}

	// TODO deserialize struct from json file with id in name

	switch(getMimetype(r)) {
	case "html": renderHtml(w, thing_single, thing)
	case "json": renderJson(w, thing)
	}
}
