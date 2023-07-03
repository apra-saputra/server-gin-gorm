package controllers

import (
	"errors"
	"net/http"
	"restapi/initializers"
	"restapi/models"
	"restapi/utils"
	"restapi/utils/customTypes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {
	userReq, _ := c.Get("user")
	limitStr := c.Query("limit")
	skipStr := c.Query("skip")

	limit, err := utils.ConverStrToInt(limitStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid query", err)
		return
	}

	skip, err := utils.ConverStrToInt(skipStr)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid query", err)
		return
	}

	var tasks []models.Task
	initializers.DB.Where("user_id = ?", userReq.(models.User).ID).Offset(int(skip)).Limit(int(limit)).Find(&tasks)

	var count int64
	initializers.DB.Model(&models.Task{}).Where("user_id = ?", userReq.(models.User).ID).Count(&count)

	response := struct {
		Payload []models.Task `json:"payload"`
		Count   int64         `json:"count"`
		Limit   int           `json:"limit"`
		Skip    int           `json:"skip"`
	}{
		Payload: tasks,
		Count:   count,
		Limit:   int(limit),
		Skip:    int(skip),
	}

	utils.JSONResponse(c, http.StatusOK, response, "Success Get Tasks", nil)
}

func CreateTask(c *gin.Context) {
	// Mendapatkan nilai query 'userid'
	userReq, _ := c.Get("user")

	// Mendapatkan data dari body request
	var request customTypes.ReqFormTask

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	var user = userReq.(models.User)

	// Membuat objek Task
	task := models.Task{
		UserID:      user.ID,
		Description: request.Description,
		Name:        request.Name,
		Date:        request.Date,
	}

	// Simpan tugas ke database
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to create task", result.Error)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, task, "Task created successfully", nil)
}

func GetTaskByIndex(c *gin.Context) {
	taskid := c.Param("id")

	var task models.Task
	result := initializers.DB.First(&task, taskid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Data Not Found", result.Error)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Success Get Task", nil)
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	var formTask customTypes.ReqFormTask

	if err := c.ShouldBindJSON(&formTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	var task models.Task

	if err := initializers.DB.First(&task, taskID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	if err := initializers.DB.Model(&task).Where("id = ?", taskID).Updates(formTask).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to update task", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Task updated successfully", nil)
}

func CompleteTask(c *gin.Context) {
	taskID := c.Param("id")

	var task models.Task

	if err := initializers.DB.First(&task, taskID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	if err := initializers.DB.Model(&task).Where("id = ?", taskID).Update("IsComplete", true).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to update task", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Task complete successfully", nil)
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	var task models.Task

	if err := initializers.DB.First(&task, taskID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Task not found", err)
		return
	}

	if err := initializers.DB.Where("id = ?", taskID).Delete(&task).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to update task", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, task, "Task deleted successfully", nil)
}
