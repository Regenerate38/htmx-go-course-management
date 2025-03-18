package handler

import (
	"fmt"
	"htmx-go-course-management/templates"
	"net/http"
)

func CoursesListHandler(w http.ResponseWriter, r *http.Request) {
	passed := processContentType("Course")

	if passed.Name == "" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	var modifiedData [][]interface{}

	for i, row := range passed.Data {
		var modifiedRow []interface{}
		fmt.Println(i)

		modifiedRow = append(modifiedRow, map[string]interface{}{
			"Index": row[0],
			"Value": row[1],
		})

		modifiedData = append(modifiedData, modifiedRow)
	}

	passed.Data = modifiedData
	fmt.Println("Passed: ", passed)
	err := templates.TMPL.ExecuteTemplate(w, "coursesList.html", passed)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
