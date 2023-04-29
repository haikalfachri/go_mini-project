package input

import (
	"mini_project/models"

	"github.com/go-playground/validator/v10"
)

type OrderInput struct {
	UserID			uint				`json:"user_id" validate:"required"`
	VehicleID		uint				`json:"vehicle_id" validate:"required"`
	TransactionID	uint				`json:"transaction_id"`
	Transaction		models.Transaction	`json:"-"`
	RentDuration	int 				`json:"rent_duration" validate:"required,min=1"`
	Status 			string 				`json:"status" gorm:"unique" validate:"required"`
}

func (u *OrderInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
