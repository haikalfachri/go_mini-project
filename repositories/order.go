package repositories

import (
	"mini_project/database"
	"mini_project/models"
	"mini_project/models/input"
)

type OrderRepositoryImp struct {
}

func InitOrderRepository() OrderRepository {
	return &OrderRepositoryImp{}
}

func (ur *OrderRepositoryImp) Create(orderInput input.OrderInput) (models.Order, error) {
	var order models.Order = models.Order{
		UserID: orderInput.UserID,
		VehicleID: orderInput.VehicleID,
		TransactionID: orderInput.TransactionID,
		RentDuration: orderInput.RentDuration,
		Status: orderInput.Status,
		// Transaction: orderInput.Transaction,
	}

	if err := database.ConnectDB().Create(&order).Error; err != nil {
		return models.Order{}, err
	}

	if err := database.ConnectDB().Last(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}

func (ur *OrderRepositoryImp) GetAll() ([]models.Order, error) {
	var orders []models.Order

	if err := database.ConnectDB().Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func (ur *OrderRepositoryImp) GetById(id string) (models.Order, error) {
	var order models.Order

	if err := database.ConnectDB().First(&order, id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (ur *OrderRepositoryImp) UpdateRating(orderInput input.OrderInput, id string) (models.Order, error) {
	order, err := ur.GetById(id)

	if err != nil {
		return models.Order{}, err
	}

	order.OrderRate = orderInput.OrderRate

	if err := database.ConnectDB().Save(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}

func (ur *OrderRepositoryImp) Update(orderInput input.OrderInput, id string) (models.Order, error) {
	order, err := ur.GetById(id)

	if err != nil {
		return models.Order{}, err
	}

	order.UserID = orderInput.UserID
	order.VehicleID = orderInput.VehicleID
	order.TransactionID = orderInput.TransactionID
	order.RentDuration = orderInput.RentDuration
	order.Status = orderInput.Status

	if err := database.ConnectDB().Save(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}

func (ur *OrderRepositoryImp) Delete(id string) error {
	order, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.ConnectDB().Delete(&order).Error; err != nil {
		return err
	}

    return nil
}
