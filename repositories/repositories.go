package repositories

import (
	"mini_project/models"
	"mini_project/models/input"
)

type UserRepository interface {
	Register(userInput input.UserInput) (models.User, error)
	Login(userInput input.UserInput) (models.User, error)
}
