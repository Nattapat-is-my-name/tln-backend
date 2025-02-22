package Usecase

import (
	"fmt"
	entities "tln-backend/Entities"
	entitiesDtos "tln-backend/Entities/dtos"
	"tln-backend/Interfaces"
)

type UserUseCase struct {
	repo Interfaces.IUserRepository
}

func NewUserUseCase(repo Interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(registerUser *entitiesDtos.RegisterRequest) error {
	var newVendor entities.Vendor

	newVendor = entities.Vendor{
		Username: registerUser.Username,
		Password: registerUser.Password,
		Email:    registerUser.Email,
	}

	return uc.repo.CreateUser(&newVendor)
}

func (uc *UserUseCase) GetUserByID(id string) (*entitiesDtos.GetUserResponse, error) {
	return uc.repo.GetUserByID(id)
}

//	func (uc *UserUseCase) UpdateUser(userID string) error {
//		err := uc.repo.UpdateUser(userID)
//		if err != nil {
//			// Log the error or handle it as necessary
//			return fmt.Errorf("use case error: %w", err)
//		}
//		return nil
//	}
func (uc *UserUseCase) DeleteUser(id string) error {

	err := uc.repo.DeleteUser(id)
	if err != nil {
		// Log the error or handle it as necessary
		return fmt.Errorf("use case error: %w", err)
	}
	return nil
}
