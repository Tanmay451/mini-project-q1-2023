package main

import (
	"net/http"

	"github.com/gorilla/mux"

	controllers "chat-room/controllers"
)

func main() {
	r := mux.NewRouter()

	// Get all forums
	r.HandleFunc("/forums", controllers.GetForums).Methods("GET")

	// Get a forum by ID
	r.HandleFunc("/forums/{id}", controllers.GetForum).Methods("GET")

	// Create a forum
	r.HandleFunc("/forums", controllers.CreateForum).Methods("POST")

	// Update a forum
	r.HandleFunc("/forums/{id}", controllers.UpdateForum).Methods("PUT")

	// Delete a forum
	r.HandleFunc("/forums/{id}", controllers.DeleteForum).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8080", r)
}
