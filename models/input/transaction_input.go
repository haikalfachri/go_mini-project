package input

import (
	"github.com/go-playground/validator/v10"
)

type TransactionInput struct {
	Name	string	`json:"name" validate:"required"`
	Data	[]byte  `json:"data" validate:"required"`
}

func (u *TransactionInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}