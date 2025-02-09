package main

import (
	"basic-golang-api/handlers"
	"basic-golang-api/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Apply logging middleware to all routes
	r.Use(middlewares.LoggingMiddleware)

	// Public routes
	r.HandleFunc("/login", handlers.GenerateToken).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBookDetail).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
	protected.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	protected.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	protected.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
