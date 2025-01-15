package handler

import (
	"htmx-go-course-management/templates"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	templates.TMPL.ExecuteTemplate(w, "home.html", nil)
}
