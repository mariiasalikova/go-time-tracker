package repositories

import (
	"database/sql"
	"errors"
	"github.com/mariiasalikova/go-time-tracker/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (user_id, name, description) VALUES ($1, $2, $3) RETURNING id`
	err := repo.DB.QueryRow(query, task.UserID, task.Name, task.Description).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepository) GetTask(id int) (*models.Task, error) {
	var task models.Task
	query := `SELECT id, user_id, name, description, created_at FROM tasks WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&task.ID, &task.UserID, &task.Name, &task.Description, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (repo *TaskRepository) UpdateTask(id int, task *models.Task) error {
	query := `UPDATE tasks SET user_id = $1, name = $2, description = $3 WHERE id = $4`
	_, err := repo.DB.Exec(query, task.UserID, task.Name, task.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepository) ListTasks() ([]models.Task, error) {
	query := `SELECT id, user_id, name, description, created_at FROM tasks`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Name, &task.Description, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
