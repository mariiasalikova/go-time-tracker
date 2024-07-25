// controllers/time_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mariiasalikova/go-time-tracker/models"
	"github.com/mariiasalikova/go-time-tracker/services"
)

type TimeController struct {
	Service *services.TimeService
}

// StartTime starts time tracking for a task
func (ctrl *TimeController) StartTime(c *gin.Context) {
	var entry models.TimeEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.StartTime = time.Now()

	if err := ctrl.Service.StartTime(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, entry)
}

// StopTime stops time tracking for a task
func (ctrl *TimeController) StopTime(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time entry ID"})
		return
	}

	if err := ctrl.Service.StopTime(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Time tracking stopped"})
}

// GetTimeEntries returns time entries for a user within a period
func (ctrl *TimeController) GetTimeEntries(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse start and end dates from query parameters
	startDate, err := time.Parse("2006-01-02", c.Query("start"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}
	endDate, err := time.Parse("2006-01-02", c.Query("end"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	entries, err := ctrl.Service.GetTimeEntries(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)
}
