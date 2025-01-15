package router

import (
	"htmx-go-course-management/handler"

	"github.com/gorilla/mux"
)

func Setuproutes(router *mux.Router) {

	router.HandleFunc("/", handler.Homepage)
	s := router.PathPrefix("/courses").Subrouter()

	s.HandleFunc("/", handler.Homepage).Methods("GET")
}
