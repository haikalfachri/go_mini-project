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

func (us *TransactionService) Update(transactionInput input.TransactionInput, id string) (models.Transaction, error){
	return us.repository.Update(transactionInput, id)
}