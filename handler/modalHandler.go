package handler

import (
	"htmx-go-course-management/templates"
	"net/http"
)

type FormComponent struct {
	Label string
	Ftype string
	key   string
}

type Form struct {
	Label      string
	FComponent []FormComponent
}

func ModalHandler(w http.ResponseWriter, r *http.Request) {
	// modalType := r.URL.Query().Get("modalType")
	var departmentComponents = [1]FormComponent{
		{
			Label: "Department Name",
			Ftype: "text",
			key:   "department_name",
		},
	}
	templates.TMPL.ExecuteTemplate(w, "modal.html", departmentComponents)
}
