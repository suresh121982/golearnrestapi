// middleware/jwt.go
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("suresh")

func AuthenticateJWT(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]

	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "This is a protected route", "user": claims})
}
