package handler

import (
	"htmx-go-course-management/templates"
	"net/http"
)

func AddPageHandler(w http.ResponseWriter, r *http.Request) {
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

	templates.TMPL.ExecuteTemplate(w, "addpage.html", arr2)
}
