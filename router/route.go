package router

import (
	"htmx-go-course-management/handler"

	"github.com/gorilla/mux"
)

func Setuproutes(router *mux.Router) {

	router.HandleFunc("/", handler.Homepage)
	router.HandleFunc("/modal", handler.ModalHandler).Methods("GET")
	router.HandleFunc("/data", handler.ReadHandler).Methods("GET")
	router.HandleFunc("/courses-list", handler.CoursesListHandler).Methods("GET")
	router.HandleFunc("/addpage", handler.AddPageHandler).Methods("GET")
	router.HandleFunc("/stats", handler.StatsHandler).Methods("GET")
	router.HandleFunc("/data/{id}", handler.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/data", handler.AddHandler).Methods("POST")
	router.HandleFunc("/data", handler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/edit/{id}", handler.EditHandler).Methods("GET")
	router.HandleFunc("/details/{id}", handler.DetailsHandler).Methods("GET")
	router.HandleFunc("/courseDetails", handler.CourseDetailsHandler).Methods("GET")
	router.HandleFunc("/exportCourseDetailsPDF", handler.ExportCourseDetailsToPDF).Methods("GET")
	router.HandleFunc("/check-login", handler.CheckLoginHandler).Methods("GET")
	router.HandleFunc("/login", handler.LoginHandler)
	router.HandleFunc("/logout", handler.LogoutHandler)

	// s := router.PathPrefix("/courses").Subrouter()
	// s.HandleFunc("/", handler.Homepage).Methods("GET")

}
