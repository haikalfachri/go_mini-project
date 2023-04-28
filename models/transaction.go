package models

import (

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Name	string	`json:"name"`
	Data	[]byte	`json:"data"`
}
