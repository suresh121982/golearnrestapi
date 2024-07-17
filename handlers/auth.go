package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"test": "password", // Example user
}

var jwtSecret = []byte("suresh")

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var tokenString string
	var err error
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		if password, ok := users[creds.Username]; ok && password == creds.Password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": creds.Username,
				"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Token valid for 7 days
			})

			tokenString, err = token.SignedString(jwtSecret)
			if err != nil {
				// Handle the error but don't use http.Error here
				// You can set err to a global error variable if necessary
				return
			}
		} else {
			// Handle unauthorized error
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}()

	wg.Wait() // Wait for the goroutine to finish

	// If token creation was successful, send the response
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		// If there was an error creating the token, return an internal server error
		http.Error(w, "Could not create token", http.StatusInternalServerError)
	}
}
