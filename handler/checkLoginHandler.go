package handler

import (
	"encoding/json"
	"htmx-go-course-management/middleware"
	"net/http"
)

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.CheckToken(r)
	isLoggedIn := err == nil && claims != nil

	data := struct {
		IsLoggedIn bool `json:"isLoggedIn"`
	}{
		IsLoggedIn: isLoggedIn,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
	}
}
