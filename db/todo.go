package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"done"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type TaskService struct {
	db *pgx.Conn
}

func NewTodoService() (*TaskService, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Failed to connect db: %v", err)
		return nil, err
	}

	return &TaskService{
		db: db,
	}, nil
}

func (t *TaskService) CreateTodo(title, description string) (*Task, error) {
	if title == "" || description == "" {
		log.Println("Failed to create Todo, because title or description is empty")
		return nil, errors.New("invalid title or description")
	}

	query := `
		INSERT INTO tasks (title, description, status, created_at, updated_at) 
		VALUES ($1, $2, 'new', NOW(), NOW())
		RETURNING id, title, description, status, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task Task
	if err := t.db.QueryRow(ctx, query, title, description).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at); err != nil {
		log.Printf("Failed to creating task: %v", err)
		return nil, err
	}

	log.Printf("Created task with ID: %d\n", task.ID)
	return &task, nil
}

func (t *TaskService) GetAllTasks() ([]Task, error) {
	query := `
		SELECT * FROM tasks ORDER BY created_at DESC
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		log.Printf("Failed to get all tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			log.Printf("Failed to scan: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskService) Update(id int, title, description string) (*Task, error) {
	if title == "" || description == "" {
		log.Println("Failed to create Todo, because title or description is empty")
		return nil, errors.New("invalid title or description")
	}

	query := `
		UPDATE tasks SET title = $1, description = $2, updated_at = NOW() WHERE id = $3 RETURNING id, title, description, status, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task Task
	if err := t.db.QueryRow(ctx, query, title, description, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		log.Printf("Failed to update task")
		return nil, err
	}

	log.Printf("Update tasks with ID: %d", id)
	return &task, nil
}

func (t *TaskService) Delete(id int) error {
	query := `
	DELETE FROM tasks WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exec, err := t.db.Exec(ctx, query, id)
	if err != nil {
		log.Printf("Fasiled to delete task: %v", err)
		return err
	}

	if exec.RowsAffected() == 0 {
		log.Println("Task not found")
		return nil
	}

	return nil
}

func (t *TaskService) UpdateStatus(id int, status string) (*Task, error) {
	Status := map[string]bool{
		"new":         true,
		"in_progress": true,
		"done":        true,
	}

	if !Status[status] {
		log.Println("Invalid status: must be 'new', 'in_progress', or 'done'")
		return nil, errors.New("invalid status: must be 'new', 'in_progress', or 'done'")
	}

	query := `
		UPDATE tasks 
		SET status = $1, updated_at = NOW() 
		WHERE id = $2
		RETURNING id, title, description, status, created_at, updated_at
	`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task Task
	if err := t.db.QueryRow(ctx, query, status, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			log.Println("Task not found")
			return nil, err
		}
		log.Printf("Failed to update status: %v", err)
		return nil, err
	}

	return &task, nil
}
