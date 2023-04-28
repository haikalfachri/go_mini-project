package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	NumberPlate	string  `json:"number_plate" gorm:"unique"`
	Type		string  `json:"type" gorm:"type:enum('car', 'motorcycle');default:'car';not_null"`
	Name     	string 	`json:"name"`
	Price		float64	`json:"price"`
	Status		string  `json:"status" gorm:"type:enum('available', 'unvailable');default:'available';not_null"`
	Rating		float64 `json:"rating"`
	Orders    	[]Order	`json:"orders" gorm:"foreignKey:VehicleID"`
}