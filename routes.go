package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.NotFoundHandler = FileHandler{Path:"404.html"}

	return router
}

type NotFoundHandler struct {
	Path string
}

func (f NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "404.html")
}

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		Index,
	},
	Route{
		"things",
		"GET",
		"/list",
		List,
	},
	Route{
		"thing",
		"GET",
		"/show/{thingId}",
		View,
	},
	Route{
		"edit",
		"GET",
		"/edit/{thingId}",
		EditForm,
	},
	Route{
		"edit",
		"POST",
		"/edit/{thingId}",
		EditStore,
	},
	Route{
		"new",
		"GET",
		"/new",
		NewForm,
	},
	Route{
		"new",
		"POST",
		"/new",
		NewStore,
	},
}
