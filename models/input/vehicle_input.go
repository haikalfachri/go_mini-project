package input

import "github.com/go-playground/validator/v10"

type VehicleInput struct {
	NumberPlate string 	`json:"number_plate" gorm:"unique" validate:"required"`
	Type		string 	`json:"type" validate:"required"`
	Name     	string 	`json:"name" validate:"required"`
	Price		float64 `json:"price" validate:"required"`
	Rating		float64 `json:"rating" validate:"min=0,max=5"`
}

func (u *VehicleInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

