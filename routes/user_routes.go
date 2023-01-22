package routes

import (
	"github.com/Funmi4194/funmimod/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(route *mux.Router) {
	//two endpoints
	route.HandleFunc("/register/", controllers.CreateAccount()).Methods("POST")
	route.HandleFunc("/user/{Name}", controllers.GetUserDetails()).Methods("GET")
}
