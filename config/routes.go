package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/goweb4/handlers"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"AddUser",
		"POST",
		"/",
		handler.AddUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/",
		handler.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/{id}",
		handler.DeleteUser,
	},
	Route{
		"AddOrder",
		"GET",
		"/",
		orderHandler.AddOrder,
	},
	Route{
		"EditOrder",
		"GET",
		"/{id}",
		orderHandler.EditOrder,
	},
	Route{
		"UpdateOrder",
		"PUT",
		"/{id}",
		orderHandler.UpdateOrder,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
