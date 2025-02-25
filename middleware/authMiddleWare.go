package middleware

import (
	"fmt"
	"htmx-go-course-management/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// CheckToken parses the token from the request's Authorization header or cookies
func CheckToken(r *http.Request) (*jwt.StandardClaims, error) {
	var tokenString string

	// Look for token in Authorization header
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = authHeader[7:] // Extract token from Authorization header
	} else {
		// Alternatively, check for the token in cookies
		cookie, err := r.Cookie("jwt")
		if err == nil {
			tokenString = cookie.Value
		}
	}

	// If no token is found
	if tokenString == "" {
		return nil, fmt.Errorf("no token found")
	}

	// Parse the token using your utility function
	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
