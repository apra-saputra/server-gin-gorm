package controller

import (
	"net/http"
	"restapi/initializers"
	"restapi/modules/model"
	"restapi/modules/repository"
	"restapi/modules/response"
	"restapi/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCalenderWithTask(c *gin.Context) {
	bySearch := c.Param("search")

	if bySearch == "year" {
		GetTasksByYear(c)
	} else if bySearch == "month" {
		GetTaskByMonth(c)
	} else {
		getTaskByDay(c)
	}
}

func GetTasksByYear(c *gin.Context) {
	userReq, _ := c.Get("user")
	yearStr := c.Query("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid year", err)
		return
	}

	payload := struct {
		Tasks map[string][]response.TaskResponse `json:"tasks"`
		Year  int                                `json:"year"`
	}{
		Tasks: make(map[string][]response.TaskResponse),
		Year:  year,
	}

	for i := 1; i <= 12; i++ {
		month := time.Month(i)
		monthName := month.String()

		payload.Tasks[monthName] = make([]response.TaskResponse, 0)
	}

	startYear := year
	endYear := year + 1

	startDate := time.Date(startYear, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, 1, 1, 0, 0, 0, 0, time.UTC)

	taskRepo := repository.NewTaskRepository(initializers.DB)

	tasks, err := taskRepo.FindAllByDateRange(userReq.(model.User).ID, startDate, endDate)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to get tasks", err)
		return
	}

	for _, task := range tasks {

		month := task.Date.Month().String()
		payload.Tasks[month] = append(payload.Tasks[month], task)
	}

	utils.JSONResponse(c, http.StatusOK, payload, "Success Get Tasks by Year", nil)
}

func GetTaskByMonth(c *gin.Context) {
	userReq, _ := c.Get("user")
	yearStr := c.Query("year")
	monthStr := c.Query("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid year", err)
		return
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid month", err)
		return
	}

	if year == 0 {
		year = time.Now().Year()
	}

	if month == 0 {
		month = int(time.Now().Month())
	}

	response := make(map[int][]struct{})

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	taskRepo := repository.NewTaskRepository(initializers.DB)

	tasks, err := taskRepo.FindAllByDateRange(userReq.(model.User).ID, startDate, endDate)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to get tasks", err)
		return
	}

	for _, task := range tasks {
		day := task.Date.Day()
		response[day] = append(response[day], struct{}{})
	}

	utils.JSONResponse(c, http.StatusOK, response, "Success Get Tasks by Month", nil)
}

func getTaskByDay(c *gin.Context) {

}
