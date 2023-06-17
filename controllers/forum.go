package controllers

import (
	models "chat-room/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetForums(w http.ResponseWriter, r *http.Request) {
	// Get all forums from the database
	forums, _ := models.GetAllForums()

	// Marshal the forums to JSON and write them to the response writer
	fmt.Println(forums)
	json.NewEncoder(w).Encode(forums)
}

func GetForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request

	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get the forum from the database
	forum, _ := models.GetForumByID(forumID)

	// Marshal the forum to JSON and write it to the response writer
	json.NewEncoder(w).Encode(forum)
}

func CreateForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum data from the request body
	forumData := models.Forum{}
	err := json.NewDecoder(r.Body).Decode(&forumData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the forum in the database
	forumID, _ := models.CreateForumDB(forumData)

	// Redirect the user to the forum page
	http.Redirect(w, r, fmt.Sprintf("/forums/%d", forumID), http.StatusFound)
}

func UpdateForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request
	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get the forum data from the request body
	forumData := models.Forum{}
	err := json.NewDecoder(r.Body).Decode(&forumData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the forum in the database
	models.UpdateForumDB(forumID, forumData)

	// Redirect the user to the forum page
	http.Redirect(w, r, fmt.Sprintf("/forums/%d", forumID), http.StatusFound)
}

func DeleteForum(w http.ResponseWriter, r *http.Request) {
	// Get the forum ID from the request
	forumID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Delete the forum from the database
	models.DeleteForumDB(forumID)

	// Redirect the user to the forums page
	http.Redirect(w, r, "/forums", http.StatusFound)
}
