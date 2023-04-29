package repositories

import (
	"mini_project/models"
	"mini_project/models/input"
)

type UserRepository interface {
	Register(userInput input.UserInput) (models.User, error)
	Login(userInput input.UserInput) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	Update(userInput input.UserInput, id string) (models.User, error)
	Delete(id string) error
}

type VehicleRepository interface {
	Create(vehicleInput input.VehicleInput) (models.Vehicle, error)
	GetByName(vehicleInput input.VehicleInput) ([]models.Vehicle, error) 
	GetAll() ([]models.Vehicle, error)
	GetById(id string) (models.Vehicle, error)
	Update(vehicleInput input.VehicleInput, id string) (models.Vehicle, error)
	Delete(id string) error
}