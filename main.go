package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Funmi4194/funmimod/configs"
	"github.com/Funmi4194/funmimod/routes"
	"github.com/gorilla/mux"
)

func main() {
	//create a route handler
	h := mux.NewRouter()
	//run database
	configs.ConnectDB()
	//To run routes
	routes.UserRoute(h)
	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	fmt.Printf("Starting server on http://localhost:%d\n", 8080)
	log.Fatal(s.ListenAndServe())
}
