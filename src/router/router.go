package router

import (
	"github.com/gorilla/mux"
)

// Return a router with all the configured routes
func Generate() *mux.Router {
	return mux.NewRouter()

}
