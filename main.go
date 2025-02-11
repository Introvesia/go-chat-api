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

	// Apply logging middleware to all routes
	r.Use(middlewares.LoggingMiddleware)

	// Public routes
	r.HandleFunc("/auth/login", handlers.Login).Methods("PUT")
	r.HandleFunc("/auth/register", handlers.Register).Methods("POST")
	r.HandleFunc("/auth/forgot-password", handlers.ForgotPassword).Methods("PUT")
	r.HandleFunc("/auth/reset-password", handlers.ResetPassword).Methods("PUT")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middlewares.AuthMiddleware)
	protected.HandleFunc("/channels", handlers.GetChannels).Methods("GET")
	protected.HandleFunc("/channels/{id}/members", handlers.GetChannelMembers).Methods("GET")

	fmt.Printf("Server running on port %s...\n", libs.GetEnv("PORT"))
	err := http.ListenAndServe(":"+libs.GetEnv("PORT"), r)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
