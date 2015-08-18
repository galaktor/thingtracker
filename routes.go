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

	return router
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
		"/things",
		Things,
	},
	Route{
		"thing",
		"GET",
		"/things/{thingId}",
		SingleThing,
	},
	Route{
		"reload",
		"GET",
		"/reload",
		Reload,
	},
}
