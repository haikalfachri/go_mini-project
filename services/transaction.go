package services

import (
	"mini_project/repositories"
	"mini_project/models"
	"mini_project/models/input"
	"mini_project/middlewares"
)

type TransactionService struct {
	repository repositories.TransactionRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitTransactionService(jwtAuth *middlewares.JWTConfig) TransactionService {
	return TransactionService{
		repository: &repositories.TransactionRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *TransactionService) Create(transactionInput input.TransactionInput) (models.Transaction, error){
	return us.repository.Create(transactionInput)
}

func (us *TransactionService) GetAll() ([]models.Transaction, error){
	return us.repository.GetAll()
}

func (us *TransactionService) GetById(id string) (models.Transaction, error){
	return us.repository.GetById(id)
}

func (us *TransactionService) Update(transactionInput input.TransactionInput, id string) (models.Transaction, error){
	return us.repository.Update(transactionInput, id)
}

func (us *TransactionService) Delete(id string) (error){
	return us.repository.Delete(id)
}