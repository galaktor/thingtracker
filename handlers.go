package main

import (
//	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var list = `<html>
  <body>
    <ul>
    foo
    {{range $i, $t := .}}
      <li>Id: {{$t.Id}}, Title: {{$t.Title}}</li>
    {{end}}
    </ul>
  </body>
</html>`

var one = `<html>
  <body>
    <p>Id: {{.Id}}</p>
    <p>Title: {{.Title}}</p>
    <p>Description: {{.Description}}</p>
    <p>Due: {{.Due}}</p>
    <p>ThingName: {{.ThingName}}</p>
    <p>ThingLink: {{.ThingLink}}</p>
  </body>
</html>`


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Things(w http.ResponseWriter, r *http.Request) {
	things := []Thing{
		Thing{Id: "1", Title: "one", ThingName: "Foo", ThingLink: "http://www.foo.bar"},
		Thing{Id: "2", Title: "two", ThingName: "Bar", ThingLink: "http://www.bar.foo"},
	}

	tmpl := template.New("things")
	var err error

	if tmpl, err = tmpl.Parse(list); err != nil {
		http.Error(w, err.Error(), 500)
	}
	tmpl.Execute(w, things)

/*	if err := json.NewEncoder(w).Encode(things); err != nil {
		panic(err)
	}*/
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

	tmpl := template.New("thing")
	var err error

	if tmpl, err = tmpl.Parse(one); err != nil {
		http.Error(w, err.Error(), 500)
	}
	tmpl.Execute(w, thing)
}
