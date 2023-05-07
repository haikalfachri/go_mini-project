package database

import (
	"errors"
	"fmt"
	"log"
	"mini_project/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}
  
  
func ConnectDB() *gorm.DB {
	config := Config{
	  DB_Username: "root",
	  DB_Password: "",
	  DB_Port:     "3306",
	  DB_Host:     "localhost",
	  DB_Name:     "mini_project_db",
	}
  
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  config.DB_Username,
	  config.DB_Password,
	  config.DB_Host,
	  config.DB_Port,
	  config.DB_Name,
	)
  
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
	  panic(err)
	}

	log.Printf("successfully connected to database\n")

	return db
}

func MigrateDB(db *gorm.DB) {

	err := db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Vehicle{}, &models.Order{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

	log.Printf("successfully database migration\n")
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}

func SeedUser(db *gorm.DB) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name: "test",
		Email: "test@test.com",
		Password: string(hashedPass),
		Role: "admin",
	}

	result := db.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SeedVehicle(db *gorm.DB) (models.Vehicle, error) {
	var vehicle models.Vehicle = models.Vehicle{
		NumberPlate: "test",
		Type: "motorcycle",
		Name: "test",
		Price: 150000.00,
		Rating: 0.00,
	}

	result := db.Create(&vehicle)

	if err := result.Error; err != nil {
		return models.Vehicle{}, err
	}

	if err := result.Last(&vehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

	return vehicle, nil
}

func SeedOrder(db *gorm.DB) (models.Order, error) {
	user, err := SeedUser(db)

	if err != nil {
		return models.Order{}, err
	}

	vehicle, err := SeedVehicle(db)

	if err != nil {
		return models.Order{}, err
	}

	transaction := models.Transaction{
		Name: "test.img",
		Data: []byte{},
	}

	result := db.Create(&transaction)

	if err := result.Error; err != nil {
		return models.Order{}, err
	}

	if err := result.Last(&transaction).Error; err != nil {
		return models.Order{}, err
	}

	var order models.Order = models.Order{
		UserID: user.ID,
		VehicleID: vehicle.ID,
		TransactionID: transaction.ID,
		Transaction: transaction,
		RentDuration: 1,
		Status: "pending",
		OrderRate: 0.0,
	}

	result = db.Create(&order)

	if err := result.Error; err != nil {
		return models.Order{}, err
	}

	if err := result.Last(&order).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func CleanSeeders(db *gorm.DB) error {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	userErr := db.Exec("DELETE FROM users").Error
	vehicleErr := db.Exec("DELETE FROM vehicles").Error
	orderErr := db.Exec("DELETE FROM orders").Error
	transactionErr := db.Exec("DELETE FROM transactions").Error

	var isFailed bool = userErr != nil || vehicleErr != nil || orderErr != nil || transactionErr != nil

	if isFailed {
		return errors.New("cleaning failed")
	}

	log.Println("seeders are cleaned up successfully")

	return nil
}