package services

import (
	"mini_project/repositories"
	"mini_project/models"
	"mini_project/models/input"
	"mini_project/middlewares"
)

type UserService struct {
	repository repositories.UserRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		repository: &repositories.UserRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *UserService) Register(userInput input.UserInput) (models.User, error){
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput input.UserInput) (string, error) {
	user, err := us.repository.Login(userInput)
	if err != nil {
		return "", err
	}

	token, err := us.jwtAuth.GenerateToken(int(user.ID), user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetAll() ([]models.User, error){
	return us.repository.GetAll()
}

func (us *UserService) GetById(id string) (models.User, error){
	return us.repository.GetById(id)
}

func (us *UserService) Update(userInput input.UserInput, id string) (models.User, error){
	return us.repository.Update(userInput, id)
}

func (us *UserService) Delete(id string) (error){
	return us.repository.Delete(id)
}

