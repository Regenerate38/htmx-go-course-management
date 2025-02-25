package handler

import (
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/model"
	"reflect"
	"strconv"
	"time"

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
	case "Teacher":
		data = &model.Teacher{}
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
	case *model.Teacher:
		database.GenericAdd(database.DB, "teacher", d)

	default:
		http.Error(w, "Unsupported data type", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received data: %+v\n", data)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))

}

func populateStruct(data interface{}, formValues map[string][]string) error {
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	for key, values := range formValues {
		var field reflect.Value
		var fieldFound bool

		// Iterate through struct fields to find a matching tag or name
		for i := 0; i < typ.NumField(); i++ {
			structField := typ.Field(i)
			tag := structField.Tag.Get("db")           // Get the "db" tag
			if tag == key || structField.Name == key { // Match by tag or field name
				field = val.Field(i)
				fieldFound = true
				break
			}
		}

		if !fieldFound || !field.IsValid() || !field.CanSet() {
			fmt.Printf("Field '%s' is invalid or cannot be set\n", key)
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if len(values) > 0 {
				field.SetString(values[0])
			}
		case reflect.Int:
			if len(values) > 0 {
				intValue, err := strconv.Atoi(values[0])
				if err == nil {
					field.SetInt(int64(intValue))
				}
			}
		case reflect.Float32:
			if len(values) > 0 {
				floatValue, err := strconv.ParseFloat(values[0], 32)
				if err == nil {
					field.SetFloat(float64(floatValue))
				}
			}
		case reflect.Float64:
			if len(values) > 0 {
				floatValue, err := strconv.ParseFloat(values[0], 64)
				if err == nil {
					field.SetFloat(floatValue)
				}
			}
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) && len(values) > 0 {
				fmt.Println("Reached here for time")
				dateValue, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", values[0]) // Custom layout for your time format
				fmt.Println("aako date", values[0])
				if err == nil {
					field.Set(reflect.ValueOf(dateValue))
					fmt.Println("time set according to populate struct", field)
				} else {
					return fmt.Errorf("failed to parse time for field '%s': %v", key, err)
				}
			}
		default:
			return fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}
	return nil
}
