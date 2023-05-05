package repositories

import (
	"mini_project/database"
	"mini_project/models"
	"mini_project/models/input"
)

type TransactionRepositoryImp struct {
}

func InitTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImp{}
}

func (ur *TransactionRepositoryImp) GetById(id string) (models.Transaction, error) {
	var transaction models.Transaction

	if err := database.ConnectDB().First(&transaction, id).Error; err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (ur *TransactionRepositoryImp) Update(transactionInput input.TransactionInput, id string) (models.Transaction, error) {
	transaction, err := ur.GetById(id)

	if err != nil {
		return models.Transaction{}, err
	}

	transaction.Name = transactionInput.Name
	transaction.Data = transactionInput.Data

	if err := database.ConnectDB().Save(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}

    return transaction, nil
}