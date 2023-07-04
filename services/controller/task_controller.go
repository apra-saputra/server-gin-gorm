package controller

import (
	"errors"
	"net/http"
	"restapi/initializers"
	"restapi/services/model"
	"restapi/services/repository"
	"restapi/services/request"
	"restapi/services/response"
	"restapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {
	userReq, _ := c.Get("user")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Query limit is required", err)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Query page is required", err)
		return
	}

	taskRepo := repository.NewTaskRepository(initializers.DB)

	tasks, count, _ := taskRepo.FindAndCountAll(userReq.(model.User).ID, limit, page)

	response := struct {
		Payload []response.TaskResponse `json:"payload"`
		Count   int64                   `json:"count"`
		Limit   int                     `json:"limit"`
		Page    int                     `json:"page"`
	}{
		Payload: tasks,
		Count:   count,
		Limit:   int(limit),
		Page:    int(page),
	}

	utils.JSONResponse(c, http.StatusOK, response, "Success Get Tasks", nil)
}

func CreateTask(c *gin.Context) {
	// Mendapatkan nilai query 'userid'
	userReq, _ := c.Get("user")

	// Mendapatkan data dari body request
	var request request.FormTask

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	var user = userReq.(model.User)

	// Membuat objek Task
	task := model.Task{
		UserID:      user.ID,
		Description: request.Description,
		Name:        request.Name,
		Date:        request.Date,
	}

	// Simpan tugas ke database
	taskRepo := repository.NewTaskRepository(initializers.DB)

	taskResponse, err := taskRepo.Create(task)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to create task", err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, taskResponse, "Task created successfully", nil)
}

func GetTaskByIndex(c *gin.Context) {
	taskId := c.Param("id")

	taskRepo := repository.NewTaskRepository(initializers.DB)

	task, err := taskRepo.FindById(taskId)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Data Not Found", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Success Get Task", nil)
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")

	var formTask request.FormTask

	if err := c.ShouldBindJSON(&formTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	taskRepo := repository.NewTaskRepository(initializers.DB)

	task, err := taskRepo.UpdateTask(taskId, formTask)

	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Task updated successfully", nil)
}

func CompleteTask(c *gin.Context) {
	taskId := c.Param("id")

	taskRepo := repository.NewTaskRepository(initializers.DB)

	task, err := taskRepo.ChangeStatusComplete(taskId)

	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Task complete successfully", nil)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	taskRepo := repository.NewTaskRepository(initializers.DB)

	err := taskRepo.Destroy(taskId)

	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	id, _ := strconv.Atoi(taskId)

	response := struct {
		Id int
	}{
		Id: id,
	}

	utils.JSONResponse(c, http.StatusOK, response, "Task deleted successfully", nil)
}
