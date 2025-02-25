package handler

import (
	// "fmt"
	"htmx-go-course-management/utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		userID := r.FormValue("userID")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		token, err := utils.GenerateJWT(userID)
		if err != nil {
			http.Error(w, "Unable to generate token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "templates/login.html")
}
