package handler

import (
	"htmx-go-course-management/model"
	"htmx-go-course-management/templates"
	"net/http"
	"reflect"
	"time"
)

type FormComponent struct {
	Label string
	Ftype string
}

type Form struct {
	Label      string
	FComponent []FormComponent
}

func ModalHandler(w http.ResponseWriter, r *http.Request) {
	modalType := r.URL.Query().Get("modalType")
	var formComponents []FormComponent
	var form Form
	switch modalType {
	case "Department":
		component := model.Department{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Department",
			FComponent: formComponents,
		}
	case "Course":
		component := model.Course{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Course",
			FComponent: formComponents,
		}
	case "Chapter":
		component := model.Chapter{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Chapter",
			FComponent: formComponents,
		}
	case "Topic":
		component := model.Topic{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Topic",
			FComponent: formComponents,
		}
	case "Lecture":
		component := model.Lecture{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Lecture",
			FComponent: formComponents,
		}
	case "Resource":
		component := model.Resource{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Resource",
			FComponent: formComponents,
		}
	case "Section":
		component := model.Section{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Section",
			FComponent: formComponents,
		}
	case "Faculty":
		component := model.Faculty{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Faculty",
			FComponent: formComponents,
		}
	case "Quiz":
		component := model.Quiz{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Quiz",
			FComponent: formComponents,
		}
	case "Assignment":
		component := model.Assignment{}
		formComponents = GenerateFormComponents(component)
		form = Form{
			Label:      "Assignment",
			FComponent: formComponents,
		}
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	templates.TMPL.ExecuteTemplate(w, "modal.html", form)
}

func getInputType(fieldType reflect.Type) string {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Float32, reflect.Float64:
		return "number"
	case reflect.String:
		return "text"
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			return "datetime-local"
		}
	default:
		return "text"
	}
	return "text"
}

func GenerateFormComponents(data interface{}) []FormComponent {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Struct {
		panic("Expected a struct")
	}

	var components []FormComponent
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		label := field.Tag.Get("form")
		if label == "" {
			label = field.Name
		}
		fType := getInputType(field.Type)
		component := FormComponent{
			Label: label,
			Ftype: fType,
		}
		components = append(components, component)
	}

	return components
}
