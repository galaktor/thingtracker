package main

import(
	"html/template"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

var thing_list *template.Template
var thing_single *template.Template
var thing_edit *template.Template

func init() {
	loadTemplates()
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

	file, err = ioutil.ReadFile("thing_edit.tmpl")
	guard(err)
	thing_edit, err = template.New("thing_edit").Parse(string(file))
	guard(err)
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
