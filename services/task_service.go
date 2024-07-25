package services

import (
	"github.com/mariiasalikova/go-time-tracker/models"
	"github.com/mariiasalikova/go-time-tracker/repositories"
)

type TaskService struct {
	Repo *repositories.TaskRepository
}

func (srv *TaskService) CreateTask(task *models.Task) error {
	return srv.Repo.CreateTask(task)
}

func (srv *TaskService) GetTask(id int) (*models.Task, error) {
	return srv.Repo.GetTask(id)
}

func (srv *TaskService) UpdateTask(id int, task *models.Task) error {
	return srv.Repo.UpdateTask(id, task)
}

func (srv *TaskService) DeleteTask(id int) error {
	return srv.Repo.DeleteTask(id)
}

func (srv *TaskService) ListTasks() ([]models.Task, error) {
	return srv.Repo.ListTasks()
}
