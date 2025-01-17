package handler

import (
	"htmx-go-course-management/templates"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	arr2 := [10]string{"Course",
		"Chapter",
		"Topic",
		"Lecture",
		"Resource",
		"Section",
		"Faculty",
		"Department",
		"Quiz",
		"Assignment"}

	templates.TMPL.ExecuteTemplate(w, "home.html", arr2)
}
