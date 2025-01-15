package main

import (
	// "htmx-go-course-management/database"
	"htmx-go-course-management/router"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// database.InitDB()
	Router := mux.NewRouter()
	router.Setuproutes(Router)

	http.ListenAndServe(":4000", Router)

}
