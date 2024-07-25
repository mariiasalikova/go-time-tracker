package services

import (
	"github.com/mariiasalikova/go-time-tracker/models"
	"github.com/mariiasalikova/go-time-tracker/repositories"
	"github.com/mariiasalikova/go-time-tracker/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (srv *UserService) CreateUser(user *models.User) error {
	info, err := utils.GetPeopleInfo(user.PassportNumber[:4], user.PassportNumber[5:])
	if err != nil {
		return err
	}

	user.Surname = info.Surname
	user.Name = info.Name
	user.Patronymic = info.Patronymic
	user.Address = info.Address

	return srv.Repo.CreateUser(user)
}

func (srv *UserService) GetUser(id int) (*models.User, error) {
	return srv.Repo.GetUser(id)
}

func (srv *UserService) UpdateUser(id int, user *models.User) error {
	return srv.Repo.UpdateUser(id, user)
}

func (srv *UserService) DeleteUser(id int) error {
	return srv.Repo.DeleteUser(id)
}

func (srv *UserService) ListUsers() ([]models.User, error) {
	return srv.Repo.ListUsers()
}
