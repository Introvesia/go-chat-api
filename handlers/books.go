package handlers

import (
	"basic-golang-api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory book storage
var books = map[string]models.Book{
	"1": {ID: "1", Title: "1984", Author: "George Orwell"},
	"2": {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get all books
func GetBookDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Check if book exists
	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Decode request body into book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = id // Ensure ID remains unchanged

	// Update the book in the map
	books[id] = book

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if ID is provided
	if newBook.ID == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	// Check if the book ID already exists
	if _, exists := books[newBook.ID]; exists {
		http.Error(w, "Book ID already exists", http.StatusConflict)
		return
	}

	// Store the new book in the map
	books[newBook.ID] = newBook

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// Update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Check if book exists
	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Decode request body into book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = id // Ensure ID remains unchanged

	// Update the book in the map
	books[id] = book

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id) // Delete book from the map
	w.WriteHeader(http.StatusNoContent)
}
