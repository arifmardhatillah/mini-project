package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateAdmin(admin *model.User) error {
	if err := config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() (users []model.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *model.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *model.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginAdmin(admin *model.User) error {
	if err := config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error; err != nil {
		return err
	}
	return nil
}
