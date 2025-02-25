package handler

import (
	"encoding/json"
	"fmt"
	"htmx-go-course-management/database"

	"net/http"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	var err error
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
	case "Course":
		err = database.DeleteCourse(database.DB, id)
	case "Chapter":
		err = database.DeleteChapter(database.DB, id)
	case "Topic":
		err = database.DeleteTopic(database.DB, id)
	case "Lecture":
		err = database.DeleteLecture(database.DB, id)
	case "Resource":
		err = database.DeleteResource(database.DB, id)
	case "Section":
		err = database.DeleteSection(database.DB, id)
	case "Faculty":
		err = database.DeleteFaculty(database.DB, id)
	case "Quiz":
		err = database.DeleteQuiz(database.DB, id)
	case "Teacher":
		err = database.DeleteTeacher(database.DB, id)
	case "Assignment":
		err = database.DeleteAssignment(database.DB, id)
	default:
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}
	if err != nil {
		message = fmt.Sprintf("%s %d", err.Error(), http.StatusInternalServerError)
		fmt.Println(message)
	}
	if err == nil {
		message = "Delete operation successfully done"
	}

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
