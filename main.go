// main.go
package main

import (
	"net/http"

	"examples.com/webapp/handlers"
	"examples.com/webapp/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/protected", middleware.AuthenticateJWT).Methods("GET")

	http.ListenAndServe(":8080", r)
}
