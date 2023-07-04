package repository

import (
	"restapi/services/model"
	"restapi/services/request"
	"restapi/services/response"
	"time"
)

func (repo *TaskRepository) Create(task model.Task) (response.TaskResponse, error) {
	if err := repo.Db.Create(&task).Error; err != nil {
		return task.GetTask(), err
	}
	return task.GetTask(), nil
}

func (repo *TaskRepository) FindAndCountAll(userID uint, limit, page int) ([]response.TaskResponse, int64, error) {
	var tasks []model.Task
	var count int64

	offset := limit * page

	err := repo.Db.
		Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&tasks).
		Error
	if err != nil {
		return nil, 0, err
	}

	err = repo.Db.Model(&model.Task{}).
		Where("user_id = ?", userID).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}

	tasksResponse := model.GetTasksToResponses(tasks)

	return tasksResponse, count, nil
}

func (repo *TaskRepository) FindAllByDateRange(userID uint, startDate time.Time, endDate time.Time) ([]response.TaskResponse, error) {
	var tasks []model.Task

	if err := repo.Db.Where("user_id = ? AND date >= ? AND date < ?", userID, startDate, endDate).Find(&tasks).Error; err != nil {
		return nil, err
	}

	tasksResponse := model.GetTasksToResponses(tasks)
	return tasksResponse, nil
}

func (repo *TaskRepository) FindById(id string) (response.TaskResponse, error) {
	var task model.Task

	err := repo.Db.First(&task, id).
		Error
	if err != nil {
		return task.GetTask(), err
	}

	return task.GetTask(), nil
}

func (repo *TaskRepository) ChangeStatusComplete(id string) (response.TaskResponse, error) {
	var task model.Task

	err := repo.Db.Model(&model.Task{}).Where("id = ?", id).Update("IsComplete", true).Find(&task).Error

	if err != nil {
		return task.GetTask(), err
	}
	return task.GetTask(), nil
}

func (repo *TaskRepository) Destroy(id string) error {
	var task model.Task

	err := repo.Db.Where("id = ?", id).Delete(&task).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *TaskRepository) UpdateTask(id string, form request.FormTask) (response.TaskResponse, error) {
	var task model.Task

	if err := repo.Db.First(&task, id).Error; err != nil {
		return task.GetTask(), err
	}

	repo.Db.Model(&task).Where("id = ?", id).Updates(form)

	return task.GetTask(), nil
}
