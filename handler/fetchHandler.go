package handler

import (
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"htmx-go-course-management/templates"
	"net/http"
	"reflect"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.URL.Query().Get("contentType")
	passed := processContentType(contentType)

	if passed.Name == "" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	err := templates.TMPL.ExecuteTemplate(w, "table.html", passed)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func processContentType(contentType string) model.PassedData {
	var items interface{}
	var itemType reflect.Type
	var name string

	switch contentType {
	case "Department":
		items, _ = database.GetDept(database.DB)
		itemType = reflect.TypeOf(model.Department{})
		name = "Department"
	case "Course":
		items, _ = database.GetCourses(database.DB)
		itemType = reflect.TypeOf(model.Course{})
		name = "Courses"
	case "Chapter":
		items, _ = database.GetChapters(database.DB)
		itemType = reflect.TypeOf(model.Chapter{})
		name = "Chapters"
	case "Topic":
		items, _ = database.GetTopics(database.DB)
		itemType = reflect.TypeOf(model.Topic{})
		name = "Topics"
	case "Lecture":
		items, _ = database.GetLectures(database.DB)
		itemType = reflect.TypeOf(model.Lecture{})
		name = "Lecture"
	case "Resource":
		items, _ = database.GetResource(database.DB)
		itemType = reflect.TypeOf(model.Resource{})
		name = "Resource"
	case "Section":
		items, _ = database.GetSection(database.DB)
		itemType = reflect.TypeOf(model.Section{})
		name = "Section"
	case "Faculty":
		items, _ = database.GetFaculty(database.DB)
		itemType = reflect.TypeOf(model.Faculty{})
		name = "Faculty"
	case "Quiz":
		items, _ = database.GetQuiz(database.DB)
		itemType = reflect.TypeOf(model.Quiz{})
		name = "Quiz"
	case "Assignment":
		items, _ = database.GetAssignment(database.DB)
		itemType = reflect.TypeOf(model.Assignment{})
		name = "Assignment"
	default:
		return model.PassedData{}
	}

	header := extractHeaders(itemType)
	rows := extractRows(items)

	return model.PassedData{
		Name:   name,
		Header: header,
		Data:   rows,
	}
}

func extractHeaders(itemType reflect.Type) []string {
	var header []string
	for i := 0; i < itemType.NumField(); i++ {
		header = append(header, itemType.Field(i).Name)
	}
	return header
}

func extractRows(items interface{}) [][]interface{} {
	var rows [][]interface{}
	itemValue := reflect.ValueOf(items)

	for i := 0; i < itemValue.Len(); i++ {
		var row []interface{}
		item := itemValue.Index(i)

		for j := 0; j < item.NumField(); j++ {
			row = append(row, item.Field(j).Interface())
		}

		rows = append(rows, row)
	}
	return rows
}
