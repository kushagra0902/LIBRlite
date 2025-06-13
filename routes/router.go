package routes

import (
	"github/Kushagra0902/LIBRlite/controller"

	"github.com/gorilla/mux"
)

func RouteHandler() *mux.Router{
	r:= mux.NewRouter()
	r.HandleFunc("/", controller.AddMessage).Methods("POST")
	
	return r
}