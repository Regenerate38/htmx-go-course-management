package handler

import (
	"fmt"
	"htmx-go-course-management/database"
	"htmx-go-course-management/templates"
	"net/http"
	"strings"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	arr1 := [11]string{"Course",
		"Chapter",
		"Topic",
		"Lecture",
		"Resource",
		"Section",
		"Faculty",
		"Department",
		"Quiz",
		"Assignment",
		"Teacher"}

	type table_data struct {
		TableName string
		Count     int
	}

	arr2 := []table_data{}

	for index, element := range arr1 {
		query := fmt.Sprintf("Select COUNT(*) from %s ", strings.ToLower(element))
		fmt.Println(index)
		var cnt int

		err := database.DB.QueryRow(query).Scan(&cnt)
		if err != nil {
			fmt.Printf("Error querying table %s: %v\n", element, err)
			break
		}

		data := table_data{TableName: element, Count: cnt}
		arr2 = append(arr2, data)
	}

	fmt.Println(arr2)
	err := templates.TMPL.ExecuteTemplate(w, "stats.html", arr2)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		fmt.Println("Error rendering template:", err)
	}
}
