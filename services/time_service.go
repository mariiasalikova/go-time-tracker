package services

import (
	"time"

	"github.com/mariiasalikova/go-time-tracker/models"
	"github.com/mariiasalikova/go-time-tracker/repositories"
)

type TimeService struct {
	Repo *repositories.TimeRepository
}

func (srv *TimeService) StartTime(entry *models.TimeEntry) error {
	return srv.Repo.StartTime(entry)
}

func (srv *TimeService) StopTime(id int) error {
	return srv.Repo.StopTime(id)
}

func (srv *TimeService) GetTimeEntries(userID int, startDate, endDate time.Time) ([]models.TimeEntry, error) {
	return srv.Repo.GetTimeEntries(userID, startDate, endDate)
}
