package routes

import (
	"github/Kushagra0902/LIBRlite/controller"

	"github.com/gorilla/mux"
)

func RouteHandler() *mux.Router{
	r:= mux.NewRouter()
	r.HandleFunc("/", controller.AddMessage).Methods("POST")
	r.HandleFunc("/{ts}", controller.GetMessages).Methods("GET")
	r.HandleFunc("/", controller.GetAllMessages).Methods("GET")
	return r
}