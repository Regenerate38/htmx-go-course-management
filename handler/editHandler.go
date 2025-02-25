package handler

import (
	"fmt"
	"net/http"

	// "strconv"
	"htmx-go-course-management/database"
	// "htmx-go-course-management/model"
	"htmx-go-course-management/templates"

	"github.com/gorilla/mux"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println(id)
	contentType := r.URL.Query().Get("contentType")
	var err error
	// id, err := strconv.Atoi(idStr) // Convert ID to integer
	// if err != nil {
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }

	var item interface{}

	switch contentType {
	case "Department":
		item, err = database.GetDepartmentByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Department",
			FComponent: formComponents,
		}
	case "Course":
		item, err = database.GetCourseByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Course",
			FComponent: formComponents,
		}
	case "Chapter":
		item, err = database.GetChapterByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Chapter",
			FComponent: formComponents,
		}

	case "Topic":
		item, err = database.GetTopicByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Topic",
			FComponent: formComponents,
		}
	case "Lecture":
		item, err = database.GetLectureByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Lecture",
			FComponent: formComponents,
		}
	case "Resource":
		item, err = database.GetResourceByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Resource",
			FComponent: formComponents,
		}
	case "Section":
		item, err = database.GetSectionByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Section",
			FComponent: formComponents,
		}
	case "Faculty":
		item, err = database.GetFacultyByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Faculty",
			FComponent: formComponents,
		}
	case "Quiz":
		item, err = database.GetQuizByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Quiz",
			FComponent: formComponents,
		}
	case "Assignment":
		item, err = database.GetAssignmentByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Assignment",
			FComponent: formComponents,
		}
	case "Teacher":
		item, err = database.GetTeacherByID(database.DB, id)
		formComponents := GenerateFormValues(item)
		item = Form{
			Label:      "Teacher",
			FComponent: formComponents,
		}
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to fetch item: %v", err), http.StatusInternalServerError)
		return
	}

	err = templates.TMPL.ExecuteTemplate(w, "editModal.html", item)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
