package application

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"EstimatesPriceGet",
		"GET",
		"/v1/estimates/price",
		EstimatesPriceGet,
	},

	Route{
		"EstimatesTimeGet",
		"GET",
		"/v1/estimates/time",
		EstimatesTimeGet,
	},

	Route{
		"ProductsGet",
		"GET",
		"/v1/products",
		ProductsGet,
	},

	Route{
		"HistoryGet",
		"GET",
		"/v1/history",
		HistoryGet,
	},

	Route{
		"MeGet",
		"GET",
		"/v1/me",
		MeGet,
	},
}
