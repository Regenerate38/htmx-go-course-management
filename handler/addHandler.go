package handler

import (
	"encoding/json"
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"reflect"
	"strconv"

	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	contentType := r.FormValue("contentType")

	var data interface{}

	var message string

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
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}
	if err := populateStruct(data, r.Form); err != nil {
		http.Error(w, "Failed to populate struct", http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)

	switch d := data.(type) {
	case *model.Department:
		database.GenericAdd(database.DB, "department", d)
	case *model.Course:
		database.GenericAdd(database.DB, "course", d)
	case *model.Chapter:
		database.GenericAdd(database.DB, "chapter", d)
	case *model.Topic:
		database.GenericAdd(database.DB, "topic", d)
	case *model.Lecture:
		database.GenericAdd(database.DB, "lecture", d)
	case *model.Resource:
		database.GenericAdd(database.DB, "resource", d)
	case *model.Section:
		database.GenericAdd(database.DB, "section", d)
	case *model.Faculty:
		database.GenericAdd(database.DB, "faculty", d)
	case *model.Quiz:
		database.GenericAdd(database.DB, "quiz", d)
	case *model.Assignment:
		database.GenericAdd(database.DB, "assignment", d)

	default:
		http.Error(w, "Unsupported data type", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received data: %+v\n", data)

	response := map[string]string{
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func populateStruct(data interface{}, formValues map[string][]string) error {
	val := reflect.ValueOf(data).Elem()
	for key, values := range formValues {
		field := val.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(values[0])
			case reflect.Int:
				intValue, err := strconv.Atoi(values[0])
				if err == nil {
					field.SetInt(int64(intValue))
				}
			case reflect.Float32:
				floatValue, err := strconv.ParseFloat(values[0], 32)
				if err == nil {
					field.SetFloat(float64(floatValue))
				}
			case reflect.Float64:
				floatValue, err := strconv.ParseFloat(values[0], 64)
				if err == nil {
					field.SetFloat(floatValue)
				}
			default:
				return fmt.Errorf("unsupported field type: %s", field.Kind())
			}
		}
	}
	return nil
}
