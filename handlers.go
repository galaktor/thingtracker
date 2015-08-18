package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"io/ioutil"
	"path/filepath"
	"io"

	"github.com/gorilla/mux"
)

var thing_list *template.Template
var thing_single *template.Template

var things []Thing

func init() {
	loadTemplates()
	t, err := getThings()
	guard(err)
	things = t
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

func getThings() ([]Thing, error) {
	files, err := ioutil.ReadDir("store")
	guard(err)

	out := []Thing{}
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == "*.thing" {
				out = append(out, Thing{})
			}
		}
	}
	
	return out, err	
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Things(w http.ResponseWriter, r *http.Request) {
	things := []Thing{
		Thing{Id: "1", Title: "one", ThingName: "Foo", ThingLink: "http://www.foo.bar"},
		Thing{Id: "2", Title: "two", ThingName: "Bar", ThingLink: "http://www.bar.foo"},
	}

	renderJson(w, things)
}

func renderJson(out io.Writer, i interface{}) error {
	if err := json.NewEncoder(out).Encode(i); err != nil {
		return err
	}

	return nil
}

func renderHtml(out io.Writer, t *template.Template, i interface{}) error {
	if err := t.Execute(out, i); err != nil {
		return err
	}

	return nil
}

func getMimetype(r *http.Request) string {
	mimetype := mux.Vars(r)["mimetype"]
	if mimetype == "" {
		mimetype = "html"
	}

	return mimetype
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
