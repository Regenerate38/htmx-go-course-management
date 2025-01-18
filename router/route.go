package router

import (
	"htmx-go-course-management/handler"

	"github.com/gorilla/mux"
)

func Setuproutes(router *mux.Router) {

	router.HandleFunc("/", handler.Homepage)
	router.HandleFunc("/modal", handler.ModalHandler).Methods("GET")
	router.HandleFunc("/data", handler.ReadHandler).Methods("GET")
	router.HandleFunc("/addpage", handler.AddPageHandler).Methods("GET")
	router.HandleFunc("/data/{id}", handler.DeleteHandler).Methods("DELETE")
	// router.HandleFunc("/data/{id}", handler.UpdateHandler).Methods("PUT", "POST")

	// s := router.PathPrefix("/courses").Subrouter()
	// s.HandleFunc("/", handler.Homepage).Methods("GET")

}
