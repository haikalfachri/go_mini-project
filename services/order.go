package services

import (
	"mini_project/repositories"
	"mini_project/models"
	"mini_project/models/input"
	"mini_project/middlewares"
)

type OrderService struct {
	repository repositories.OrderRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitOrderService(jwtAuth *middlewares.JWTConfig) OrderService {
	return OrderService{
		repository: &repositories.OrderRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *OrderService) Create(orderInput input.OrderInput) (models.Order, error){
	return us.repository.Create(orderInput)
}

func (us *OrderService) GetAll() ([]models.Order, error){
	return us.repository.GetAll()
}

func (us *OrderService) GetById(id string) (models.Order, error){
	return us.repository.GetById(id)
}

func (us *OrderService) GetHistory(id string) ([]models.Order, error){
	return us.repository.GetHistory(id)
}


func (us *OrderService) UpdateStatus(id string) (models.Order, error){
	return us.repository.UpdateStatus(id)
}

func (us *OrderService) UpdateRating(orderInput input.OrderInput, id string) (models.Order, error){
	return us.repository.UpdateRating(orderInput, id)
}

func (us *OrderService) Update(orderInput input.OrderInput, id string) (models.Order, error){
	return us.repository.Update(orderInput, id)
}

func (us *OrderService) Delete(id string) (error){
	return us.repository.Delete(id)
}