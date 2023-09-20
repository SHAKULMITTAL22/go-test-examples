package main

import (
	"testing"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type testRepository struct {
	db *sql.DB
}

type TestTask struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      string
	Priority    string
	Description string
	CreatedAt   time.Time
}

func (r testRepository) Find(id int) (TestTask, error) {
	var task TestTask

	rows, _ := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()

	// If there is no record
	if rows.Next() == false {
		return task, errors.New("Record not found!")
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
		if err != nil {
			return task, errors.New("Records could not be matched!")
		}
	}

	if err := rows.Err(); err != nil {
		return task, err
	}

	return task, nil
}

func TestFind_08d58dd201(t *testing.T) {
	// TODO: Replace with your own database connection
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}

	repo := testRepository{db: db}

	// Test case 1: Existing task
	task, err := repo.Find(1) // TODO: Replace with an existing task ID
	if err != nil {
		t.Error(fmt.Sprintf("Failed to find existing task: %v", err))
	} else {
		t.Log(fmt.Sprintf("Successfully found existing task: %v", task))
	}

	// Test case 2: Non-existing task
	task, err = repo.Find(9999) // TODO: Replace with a non-existing task ID
	if err == nil || err.Error() != "Record not found!" {
		t.Error(fmt.Sprintf("Expected error for non-existing task, got: %v", err))
	} else {
		t.Log("Successfully detected non-existing task")
	}
}
