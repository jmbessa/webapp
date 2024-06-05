package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func main() {

	utils.LoadTemplates()

	r := router.Generate()

	fmt.Println("Escutando na porta 3000")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		log.Printf("Registered route: %s %s", methods, path)
		return nil
	})
	log.Fatal(http.ListenAndServe(":3000", r))
}
