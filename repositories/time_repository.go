// repositories/time_repository.go
package repositories

import (
	"database/sql"
	"github.com/mariiasalikova/go-time-tracker/models"
	"time"
)

type TimeRepository struct {
	DB *sql.DB
}

func (repo *TimeRepository) StartTime(entry *models.TimeEntry) error {
	query := `INSERT INTO time_entries (user_id, task_id, start_time) VALUES ($1, $2, $3) RETURNING id`
	err := repo.DB.QueryRow(query, entry.UserID, entry.TaskID, entry.StartTime).Scan(&entry.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TimeRepository) StopTime(id int) error {
	query := `UPDATE time_entries SET end_time = $1 WHERE id = $2`
	_, err := repo.DB.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TimeRepository) GetTimeEntries(userID int, startDate, endDate time.Time) ([]models.TimeEntry, error) {
	query := `SELECT id, user_id, task_id, start_time, end_time FROM time_entries WHERE user_id = $1 AND start_time BETWEEN $2 AND $3`
	rows, err := repo.DB.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.TimeEntry
	for rows.Next() {
		var entry models.TimeEntry
		if err := rows.Scan(&entry.ID, &entry.UserID, &entry.TaskID, &entry.StartTime, &entry.EndTime); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
