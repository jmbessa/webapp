package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all the routes of the API
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Configure puts the routes inside the router
func Configure(router *mux.Router) *mux.Router {
	routes := loginRoute

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
