package repositories

import (
	"mini_project/database"
	"mini_project/models"
	"mini_project/models/input"
	
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImp struct {
}

func InitUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (ur *UserRepositoryImp) Register(userInput input.UserInput) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost) 
	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name : userInput.Name,
		Email: userInput.Email,
		Password : string(hashedPass),
		Role: userInput.Role,
	}

	if err := database.ConnectDB().Create(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := database.ConnectDB().Last(&user).Error; err != nil {
		return models.User{}, err
	}

    return user, nil
}

func (ur *UserRepositoryImp) Login(userInput input.UserInput) (models.User, error) {
	var user models.User

	if err := database.ConnectDB().First(&user, "email = ?", userInput.Email).Error; err != nil {
		return models.User{}, err
	}

	if err :=  bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImp) GetAll() ([]models.User, error) {
	var users []models.User

	if err := database.ConnectDB().Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepositoryImp) GetById(id string) (models.User, error) {
	var user models.User

	if err := database.ConnectDB().First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepositoryImp) Update(userInput input.UserInput, id string) (models.User, error) {
	user, err := ur.GetById(id)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost) 
	if err != nil {
		return models.User{}, err
	}

	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Password = string(hashedPass)
	user.Role = userInput.Role

	if err := database.ConnectDB().Save(&user).Error; err != nil {
		return models.User{}, err
	}

    return user, nil
}

func (ur *UserRepositoryImp) Delete(id string) error {
	user, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.ConnectDB().Delete(&user).Error; err != nil {
		return err
	}

    return nil
}


