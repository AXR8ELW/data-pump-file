package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetWeights", // Name
		"GET",        // HTTP method
		"/weight",    // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}

// Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {

	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {

		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
