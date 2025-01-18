package handler

import (
	"encoding/json"
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"

	// "htmx-go-course-management/templates"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var items interface{}
	var itemType reflect.Type
	var err error
	var name string
	var message string
	vars := mux.Vars(r)
	id := vars["id"]
	var requestBody struct {
		ContentType string `json:"contentType"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	contentType := requestBody.ContentType

	fmt.Println(vars, id)

	switch contentType {
	case "Department":
		err = database.DeleteDept(database.DB, id)
		if err != nil {
			message = fmt.Sprintf("%s %s %d", w, err.Error(), http.StatusInternalServerError)
			fmt.Println(message)
		}
		items, err = database.GetDept(database.DB)
		itemType = reflect.TypeOf(model.Department{})
		name = "Department"
	case "Course":
		items, err = database.GetCourses(database.DB)
		itemType = reflect.TypeOf(model.Course{})
		name = "Courses"
	case "Chapter":
		items, err = database.GetChapters(database.DB)
		itemType = reflect.TypeOf(model.Chapter{})
		name = "Chapters"
	case "Topic":
		items, err = database.GetTopics(database.DB)
		itemType = reflect.TypeOf(model.Topic{})
		name = "Topics"
	case "Lecture":
		items, err = database.GetLectures(database.DB)
		itemType = reflect.TypeOf(model.Lecture{})
		name = "Lecture"
	case "Resource":
		items, err = database.GetResource(database.DB)
		itemType = reflect.TypeOf(model.Resource{})
		name = "Resource"
	case "Section":
		items, err = database.GetSection(database.DB)
		itemType = reflect.TypeOf(model.Section{})
		name = "Section"
	case "Faculty":
		items, err = database.GetFaculty(database.DB)
		itemType = reflect.TypeOf(model.Faculty{})
		name = "Faculty"
	case "Quiz":
		items, err = database.GetQuiz(database.DB)
		itemType = reflect.TypeOf(model.Quiz{})
		name = "Quiz"
	case "Assignment":
		items, err = database.GetAssignment(database.DB)
		itemType = reflect.TypeOf(model.Assignment{})
		name = "Assignment"
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	if err != nil {
		message = fmt.Sprintf("%s %d", err.Error(), http.StatusInternalServerError)
		fmt.Println(message)
	}

	header := extractHeaders(itemType)
	rows := extractRows(items)

	fmt.Println(header, rows, name)
	// passed := model.PassedData{
	// 	Name:   name,
	// 	Header: header,
	// 	Data:   rows,
	// }

	response := map[string]string{
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// err = templates.TMPL.ExecuteTemplate(w, "table.html", passed)
	// if err != nil {
	// 	http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	// }
}
