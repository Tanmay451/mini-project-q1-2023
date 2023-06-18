package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Forum struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetAllForums() (forums []Forum, err error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://localhost:5432/forum")
	if err != nil {
		return forums, err
	}
	// Get all forums from the database
	rows, err := db.Query("SELECT * FROM forum_chat_room ORDER BY created_at DESC")
	if err != nil {
		return forums, err
	}

	// Iterate over the rows and create forum objects
	for rows.Next() {
		var forum Forum
		err = rows.Scan(&forum.ID, &forum.Name, &forum.Description, &forum.CreatedAt, &forum.UpdatedAt)
		if err != nil {
			return forums, err
		}

		forums = append(forums, forum)
	}

	// Close the database connection
	db.Close()

	return forums, nil
}

func GetForumByID(id int) (forum Forum, err error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://localhost:5432/forum")
	if err != nil {
		return forum, err
	}

	// Get the forum from the database
	row := db.QueryRow("SELECT * FROM forum_chat_room WHERE id = $1", id)
	err = row.Scan(&forum.ID, &forum.Name, &forum.Description, &forum.CreatedAt, &forum.UpdatedAt)
	if err != nil {
		return forum, err
	}

	// Close the database connection
	db.Close()

	return forum, nil
}

func CreateForumDB(forum Forum) (id int64, err error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://localhost:5432/forum")
	if err != nil {
		return id, err
	}

	// Create the forum in the database
	stmt, err := db.Prepare("INSERT INTO forum_chat_room (name, description) VALUES ($1, $2)")
	if err != nil {
		return id, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(forum.Name, forum.Description)
	if err != nil {
		return id, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}

	// Close the database connection
	db.Close()

	return id, nil
}

func UpdateForumDB(id int, forum Forum) (err error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://localhost:5432/forum")
	if err != nil {
		return err
	}

	// Update the forum in the database
	stmt, err := db.Prepare("UPDATE forum_chat_room SET name = $1, description = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(forum.Name, forum.Description, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	// Close the database connection
	db.Close()

	return nil
}

func DeleteForumDB(id int) (err error) {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://localhost:5432/forum")
	if err != nil {
		return err
	}

	// Delete the forum from the database
	stmt, err := db.Prepare("DELETE FROM forum_chat_room WHERE id = $1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	// Close the database connection
	db.Close()

	return nil
}
