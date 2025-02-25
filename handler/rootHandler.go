package handler

import (
	"htmx-go-course-management/templates"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	arr2 := [11]string{"Course",
		"Chapter",
		"Topic",
		"Lecture",
		"Resource",
		"Section",
		"Faculty",
		"Department",
		"Quiz",
		"Assignment",
		"Teacher"}

	templates.TMPL.ExecuteTemplate(w, "home.html", arr2)
}
