package service

import (
	"cubicasa/database"
	"cubicasa/models"
)

type UserService struct {
	userRepo database.UserRepository
}

func (userService *UserService) InitService() {
	userService.userRepo.InitDB("PG")
}

func (userService *UserService) GetAllUser() ([]models.User, bool) {
	listUser, isDone := userService.userRepo.GetAllUser()
	return listUser, isDone
}

func (userService *UserService) GetUserById(id int) (models.User, bool) {
	return userService.userRepo.GetUserById(id)
}

func (userService *UserService) UpdateUser(id int, user models.User) (int, bool) {
	return userService.userRepo.UpdateUser(id, user)
}

func (userService *UserService) CreateUser(user models.User) (models.User, bool) {
	return userService.userRepo.CreateUser(user)
}

func (userService *UserService) DeleteUser(id int) (int, bool) {
	return userService.userRepo.DeleteUser(id)
}

func (userService *UserService) MoveUser(userId int, teamId int) bool {
	return userService.userRepo.MoveUser(userId, teamId)
}

func (userService *UserService) SearchUser(userName string) (map[string]string, bool) {
	return userService.userRepo.SearchUsername(userName)
}
