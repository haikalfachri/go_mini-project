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

type TransactionRepository interface {
	Create(transactionInput input.TransactionInput) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	GetById(id string) (models.Transaction, error)
	Update(transactionInput input.TransactionInput, id string) (models.Transaction, error)
	Delete(id string) error
}

type OrderRepository interface {
	Create(orderInput input.OrderInput) (models.Order, error)
	GetAll() ([]models.Order, error)
	GetById(id string) (models.Order, error)
	Update(orderInput input.OrderInput, id string) (models.Order, error)
	Delete(id string) error
}