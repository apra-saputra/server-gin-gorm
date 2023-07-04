package repository

import (
	"errors"
	"restapi/services/model"
)

func (repo *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := repo.Db.Where("email = ?", email).Find(&user).Error
	if err != nil || user.GetUser().Id == 0 {
		if err == nil {
			errorCustom := errors.New("Data not found")
			return user, errorCustom
		}
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) UpdateOtp(id uint, otp string) error {
	err := repo.Db.Model(&model.User{}).Where("id = ?", id).Update("otp", otp).Error

	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateResetOtp(id uint) error {
	err := repo.Db.Model(&model.User{}).Where("id = ?", id).Update("otp", nil).Error

	if err != nil {
		return err
	}
	return nil
}
