package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var thing_list *template.Template
var thing_single *template.Template

var things map[int]*Thing

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

func thingFileNameToId(filename string) (int, error) {
	return strconv.Atoi(strings.TrimSuffix(filename, ".thing"))
}

func deserThings() (map[int]*Thing, error) {
	files, err := ioutil.ReadDir("store")
	guard(err)

	out := make(map[int]*Thing)
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".thing" {
				id, err := thingFileNameToId(file.Name())
				guard(err)
				fullpath := "store/" + file.Name()
				content, err := ioutil.ReadFile(fullpath)
				guard(err)
				t := &Thing{}
				err = json.Unmarshal(content, t)
				guard(err)
				out[id] = t
			}
		}
	}
	
	return out, err	
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func refreshThings() (err error) {
	things, err = deserThings()
	return
}

func Things(w http.ResponseWriter, r *http.Request) {
	if err := refreshThings(); err != nil {
		http.Error(w, err.Error(), 500)
	}
	
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

func Edit(w http.ResponseWriter, r *http.Request) {
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

func NewAdd(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "ADD not implemented", 404)
}
