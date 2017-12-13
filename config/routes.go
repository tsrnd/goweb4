package config

import (
	"net/http"
	"github.com/goweb4/middleware"
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
		handlers.Index,
	},
	Route{
		"AdminIndex",
		"GET",
		"/adminIndex",
		handlers.IndexAdmin,
	},
	Route{
		"Register",
		"GET",
		"/register",
		handlers.Register,
	},
	Route{
		"Register",
		"POST",
		"/registerHandler",
		handlers.RegisterHandler,
	},
	Route{
		"UserProfile",
		"GET",
		"/userProfile",
		handlers.UserProfile,
	},
	// Route{
	// 	"CreateUser",
	// 	"POST",
	// 	"/user",
	// 	handlers.AddUser,
	// },
	Route{
		"UpdateUser",
		"PUT",
		"/user/{id}",
		handlers.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/{id}",
		handlers.DeleteUser,
	},
	// Route{
	// 	"CreateOrder",
	// 	"GET",
	// 	"/order/add",
	// 	handlers.CreateOrder,
	// },
	Route{
		"StoreOrder",
		"POST",
		"/order",
		handlers.StoreOrder,
	},
	// Route{
	// 	"EditOrder",
	// 	"GET",
	// 	"/order/{id}/edit",
	// 	handlers.EditOrder,
	// },
	// Route{
	// 	"UpdateOrder",
	// 	"PUT",
	// 	"/order/{id}",
	// 	handlers.UpdateOrder,
	// },
	Route{
		"Login",
		"GET",
		"/login",
		handlers.Login,
	},
	Route{
		"LoginHandler",
		"POST",
		"/loginHandler",
		handlers.LoginHandler,
	},
	Route{
		"LogoutHandler",
		"GET",
		"/logout",
		handlers.LogoutHandler,
	},
	Route{
		"CreateProduct",
		"GET",
		"/admin/product/add",
		middleware.FilterAdmin(handlers.CreateProduct),
	},
	Route{
		"StoreProduct",
		"POST",
		"/admin/product",
		middleware.FilterAdmin(handlers.StoreProduct),
	},
	Route{
		"CreateProduct",
		"GET",
		"/admin/product/{id:[0-9]+}",
		middleware.FilterAdmin(handlers.ShowProduct),
	},
	// Route{
	// 	"DetailProduct",
	// 	"GET",
	// 	"/detail/{id}",
	// 	handlers.DetailProduct,
	// },
	Route{
		"ShowProductGroup",
		"GET",
		"/group/{id}",
		handlers.ShowProductGroup,
	},
	Route{
		"ContactUs",
		"GET",
		"/contact",
		handlers.ContactUs,
	},
	Route{
		"AboutUs",
		"GET",
		"/about",
		handlers.AboutUs,
	},
	// Route{
	// 	"ShowCart",
	// 	"GET",
	// 	"/cart",
	// 	handlers.ShowCart,
	// },
	// Route{
	// 	"Checkout",
	// 	"GET",
	// 	"/checkout",
	// 	handlers.Checkout,
	// },
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
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
