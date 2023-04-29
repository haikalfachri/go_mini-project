package services

import (
	"mini_project/repositories"
	"mini_project/models"
	"mini_project/models/input"
	"mini_project/middlewares"
)

type VehicleService struct {
	repository repositories.VehicleRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitVehicleService(jwtAuth *middlewares.JWTConfig) VehicleService {
	return VehicleService{
		repository: &repositories.VehicleRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *VehicleService) Create(vehicleInput input.VehicleInput) (models.Vehicle, error){
	return us.repository.Create(vehicleInput)
}

func (us *VehicleService) GetByName(vehicleInput input.VehicleInput) ([]models.Vehicle, error){
	return us.repository.GetByName(vehicleInput)
}

func (us *VehicleService) GetAll() ([]models.Vehicle, error){
	return us.repository.GetAll()
}

func (us *VehicleService) GetById(id string) (models.Vehicle, error){
	return us.repository.GetById(id)
}

func (us *VehicleService) Update(vehicleInput input.VehicleInput, id string) (models.Vehicle, error){
	return us.repository.Update(vehicleInput, id)
}

func (us *VehicleService) Delete(id string) (error){
	return us.repository.Delete(id)
}