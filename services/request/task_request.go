package request

import "time"

type FormTask struct {
	Description string    `form:"description" json:"description" binding:"required"`
	Name        string    `form:"name" json:"name" binding:"required"`
	Date        time.Time `form:"date" json:"date" time_format:"2006-01-02" binding:"required"`
}
