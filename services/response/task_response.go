package response

import "time"

type TaskResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Date        time.Time `json:"date"`
}
