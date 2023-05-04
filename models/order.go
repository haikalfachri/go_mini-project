package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID         	uint    	`json:"user_id"`
	VehicleID      	uint    	`json:"vehicle_id"`
	TransactionID   uint     	`json:"transaction_id" gorm:"uniqueIndex"`
	Transaction     Transaction `json:"-" gorm:"foreignKey:TransactionID"`
	RentDuration	int  		`json:"rent_duration"`
	Status			string  	`json:"status" gorm:"type:enum('pending', 'accepted');default:'pending';not_null"`
	OrderRate		float64		`json:"order_rate"`
}