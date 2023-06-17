package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Get all forums
	r.HandleFunc("/forums", GetForums).Methods("GET")

	// Get a forum by ID
	r.HandleFunc("/forums/{id}", GetForum).Methods("GET")

	// Create a forum
	r.HandleFunc("/forums", CreateForum).Methods("POST")

	// Update a forum
	r.HandleFunc("/forums/{id}", UpdateForum).Methods("PUT")

	// Delete a forum
	r.HandleFunc("/forums/{id}", DeleteForum).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8080", r)
}

func GetForums(w http.ResponseWriter, r *http.Request) {
	// Get all forums from the database
	forums, _ := GetAllForums()

	// Marshal the forums to JSON and write them to the response writer
	fmt.Println(forums)
	json.NewEncoder(w).Encode(forums)
}

func GetForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request

	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get the forum from the database
	forum, _ := GetForumByID(forumID)

	// Marshal the forum to JSON and write it to the response writer
	json.NewEncoder(w).Encode(forum)
}

func CreateForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum data from the request body
	forumData := Forum{}
	err := json.NewDecoder(r.Body).Decode(&forumData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the forum in the database
	forumID, _ := CreateForumDB(forumData)

	// Redirect the user to the forum page
	http.Redirect(w, r, fmt.Sprintf("/forums/%d", forumID), http.StatusFound)
}

func UpdateForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request
	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get the forum data from the request body
	forumData := Forum{}
	err := json.NewDecoder(r.Body).Decode(&forumData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the forum in the database
	UpdateForumDB(forumID, forumData)

	// Redirect the user to the forum page
	http.Redirect(w, r, fmt.Sprintf("/forums/%d", forumID), http.StatusFound)
}

func DeleteForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request
	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Delete the forum from the database
	DeleteForumDB(forumID)

	// Redirect the user to the forums page
	http.Redirect(w, r, "/forums", http.StatusFound)
}
