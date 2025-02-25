package handler

import (
	"encoding/json"
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	contentType := r.FormValue("contentType")

	var data interface{}

	switch contentType {
	case "Department":
		data = &model.Department{}
	case "Course":
		data = &model.Course{}
	case "Chapter":
		data = &model.Chapter{}
	case "Topic":
		data = &model.Topic{}
	case "Lecture":
		data = &model.Lecture{}
	case "Resource":
		data = &model.Resource{}
	case "Section":
		data = &model.Section{}
	case "Faculty":
		data = &model.Faculty{}
	case "Quiz":
		data = &model.Quiz{}
	case "Assignment":
		data = &model.Assignment{}
	case "Teacher":
		data = &model.Teacher{}
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	if err := populateStruct(data, r.Form); err != nil {
		fmt.Println(r.Form)
		fmt.Println(data)
		http.Error(w, "Failed to populate struct", http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Form)
	fmt.Println(data)

	switch d := data.(type) {
	case *model.Department:
		if err := database.GenericUpdate(database.DB, "department", d); err != nil {
			http.Error(w, "Failed to update department", http.StatusInternalServerError)
			return
		}
	case *model.Course:
		if err := database.GenericUpdate(database.DB, "course", d); err != nil {
			http.Error(w, "Failed to update course", http.StatusInternalServerError)
			return
		}
	case *model.Chapter:
		if err := database.GenericUpdate(database.DB, "chapter", d); err != nil {
			http.Error(w, "Failed to update chapter", http.StatusInternalServerError)
			return
		}
	case *model.Topic:
		if err := database.GenericUpdate(database.DB, "topic", d); err != nil {
			http.Error(w, "Failed to update topic", http.StatusInternalServerError)
			return
		}
	case *model.Lecture:
		if err := database.GenericUpdate(database.DB, "lecture", d); err != nil {
			http.Error(w, "Failed to update lecture", http.StatusInternalServerError)
			return
		}
	case *model.Resource:
		if err := database.GenericUpdate(database.DB, "resource", d); err != nil {
			http.Error(w, "Failed to update resource", http.StatusInternalServerError)
			return
		}
	case *model.Section:
		if err := database.GenericUpdate(database.DB, "section", d); err != nil {
			http.Error(w, "Failed to update section", http.StatusInternalServerError)
			return
		}
	case *model.Faculty:
		if err := database.GenericUpdate(database.DB, "faculty", d); err != nil {
			http.Error(w, "Failed to update faculty", http.StatusInternalServerError)
			return
		}
	case *model.Quiz:
		if err := database.GenericUpdate(database.DB, "quiz", d); err != nil {
			http.Error(w, "Failed to update quiz", http.StatusInternalServerError)
			return
		}
	case *model.Assignment:
		if err := database.GenericUpdate(database.DB, "assignment", d); err != nil {
			http.Error(w, "Failed to update assignment", http.StatusInternalServerError)
			return
		}
	case *model.Teacher:
		if err := database.GenericUpdate(database.DB, "teacher", d); err != nil {
			http.Error(w, "Failed to update teacher", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Unsupported data type", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": fmt.Sprintf("Successfully updated %s", contentType),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
