package main

import (
	"fmt"
	"go-chat-api/handlers"
	"go-chat-api/libs"
	"go-chat-api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	libs.LoadEnv()
	r := mux.NewRouter()
	db := libs.ConnectDB()

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

	fmt.Printf("Server running on port %s...\n", libs.GetEnv("PORT"))
	err := http.ListenAndServe(":"+libs.GetEnv("PORT"), r)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
