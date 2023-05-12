package usecase

import (
	"fmt"
	"project_structure/middleware"
	"project_structure/model"
	"project_structure/model/payload"
	"project_structure/repository/database"
)

func LoginUser(user *model.User) (err error) {
	// check to db email and password
	err = database.LoginUser(user)
	if err != nil {
		fmt.Println("GetUser : Error getting user from database")
		return
	}
	// generate jwt
	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		fmt.Println("GetUser : Error Generate token")
		return
	}
	user.Token = token
	return
}

func LoginAdmin(admin *model.User) (err error) {
	// check to db email and password
	err = database.LoginAdmin(admin)
	if err != nil {
		fmt.Println("GetAdmin : Error getting admin from database")
		return
	}
	// generate jwt
	token, err := middleware.CreateToken(int(admin.ID), admin.Role)
	if err != nil {
		fmt.Println("GetUser : Error Generate token")
		return
	}
	admin.Token = token
	return
}

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Role:     "USER",
	}
	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	// generate jwt
	token, err := middleware.CreateToken(int(newUser.ID), newUser.Role)
	if err != nil {
		fmt.Println("GetUser : Error Generate token")
		return
	}
	newUser.Token = token
	err = database.UpdateUser(newUser)
	if err != nil {
		fmt.Println("UpdateUser : Error Update user")
		return
	}
	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
		Token:  newUser.Token,
	}
	return
}

func CreateAdmin(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Role:     "ADMIN",
	}
	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	// generate jwt
	token, err := middleware.CreateToken(int(newUser.ID), newUser.Role)
	if err != nil {
		fmt.Println("GetAdmin : Error Generate token")
		return
	}
	newUser.Token = token
	err = database.UpdateUser(newUser)
	if err != nil {
		fmt.Println("UpdateAdmin : Error Update admin")
		return
	}
	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
		Token:  newUser.Token,
	}
	return
}

func GetUserByEmail(email string) (*model.User, error) {
	user, err := database.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetListUsers() (users []model.User, err error) {
	users, err = database.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers : Error getting users from database")
		return
	}
	return
}

func UpdateUser(user *model.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func TopUp(userID, saldo uint) (*model.User, error) {
	user, err := database.GetUser(userID)
	if err != nil {
		return nil, err
	}
	user.Saldo += saldo
	err = database.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
