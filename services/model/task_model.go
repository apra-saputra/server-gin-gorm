package model

import (
	"restapi/services/response"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	IsComplete  bool `gorm:"default:false"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	Date        time.Time
}

func (task *Task) GetTask() response.TaskResponse {
	var statusText string

	if task.IsComplete {
		statusText = "done"
	} else {
		statusText = "not yet"
	}

	return response.TaskResponse{
		Id:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      statusText,
		Date:        task.Date,
	}
}

func GetTasksToResponses(tasks []Task) []response.TaskResponse {
	taskResponses := make([]response.TaskResponse, len(tasks))

	for i, task := range tasks {
		var statusText string

		if task.IsComplete {
			statusText = "done"
		} else {
			statusText = "not yet"
		}

		taskResponse := response.TaskResponse{
			Id:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			Status:      statusText,
			Date:        task.Date,
		}

		taskResponses[i] = taskResponse
	}

	return taskResponses
}
