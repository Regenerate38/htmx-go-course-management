package handler

// import (
// 	"htmx-go-course-management/database"
// 	"htmx-go-course-management/model"
// 	"htmx-go-course-management/templates"
// 	"net/http"
// 	"reflect"
// 	"strconv"
// )

// func ReadByIdHandler(w http.ResponseWriter, r *http.Request) {
// 	contentType := r.URL.Query().Get("contentType")
// 	id := r.URL.Query().Get("id")
// 	passed := processSpecificContentType(contentType, id)

// 	if passed.Name == "" {
// 		http.Error(w, "Invalid content type", http.StatusBadRequest)
// 		return
// 	}

// 	err := templates.TMPL.ExecuteTemplate(w, "modal.html", passed)
// 	if err != nil {
// 		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
// 	}
// }

// func processSpecificContentType(contentType string, id string) model.ModelData {
// 	var item interface{}
// 	var itemType reflect.Type
// 	var name string

// 	switch contentType {
// 	case "Department":
// 		deptID := parseID(id) // Convert id to appropriate type (int)
// 		item, _ = database.GetDepartmentByID(database.DB, deptID)
// 		itemType = reflect.TypeOf(model.Department{})
// 		name = "Department"
// 	case "Course":
// 		courseID := parseID(id)
// 		item, _ = database.GetCourseByID(database.DB, courseID)
// 		itemType = reflect.TypeOf(model.Course{})
// 		name = "Course"
// 	case "Chapter":
// 		chapterID := parseID(id)
// 		item, _ = database.GetChapterByID(database.DB, chapterID)
// 		itemType = reflect.TypeOf(model.Chapter{})
// 		name = "Chapter"
// 	case "Topic":
// 		topicID := parseID(id)
// 		item, _ = database.GetTopicByID(database.DB, topicID)
// 		itemType = reflect.TypeOf(model.Topic{})
// 		name = "Topic"
// 	case "Lecture":
// 		lectureID := parseID(id)
// 		item, _ = database.GetLectureByID(database.DB, lectureID)
// 		itemType = reflect.TypeOf(model.Lecture{})
// 		name = "Lecture"
// 	case "Resource":
// 		resourceID := parseID(id)
// 		item, _ = database.GetResourceByID(database.DB, resourceID)
// 		itemType = reflect.TypeOf(model.Resource{})
// 		name = "Resource"
// 	case "Section":
// 		sectionID := parseID(id)
// 		item, _ = database.GetSectionByID(database.DB, sectionID)
// 		itemType = reflect.TypeOf(model.Section{})
// 		name = "Section"
// 	case "Faculty":
// 		facultyID := parseID(id)
// 		item, _ = database.GetFacultyByID(database.DB, facultyID)
// 		itemType = reflect.TypeOf(model.Faculty{})
// 		name = "Faculty"
// 	case "Quiz":
// 		quizID := parseID(id)
// 		item, _ = database.GetQuizByID(database.DB, quizID)
// 		itemType = reflect.TypeOf(model.Quiz{})
// 		name = "Quiz"
// 	case "Assignment":
// 		assignID := parseID(id)
// 		item, _ = database.GetAssignmentByID(database.DB, assignID)
// 		itemType = reflect.TypeOf(model.Assignment{})
// 		name = "Assignment"
// 	default:
// 		return model.ModelData{}
// 	}

// 	header := extractHeaders(itemType)
// 	row := extractRow(item) // Extract a single row since we are getting one item

// 	return model.ModelData{
// 		Name:   name,
// 		Header: header,
// 		Data:   row,
// 	}
// }

// // Helper function to parse ID from string to int
// func parseID(id string) int {
// 	idInt, err := strconv.Atoi(id) // Convert string ID to int
// 	if err != nil {
// 		return 0 // Handle error appropriately; could return an error or zero
// 	}
// 	return idInt
// }

// func extractRow(item interface{}) []interface{} {
// 	var row []interface{}
// 	itemValue := reflect.ValueOf(item)

// 	for i := 0; i < itemValue.NumField(); i++ {
// 		row = append(row, itemValue.Field(i).Interface())
// 	}
// 	return row
// }
