package auth_service

import (
	"database/sql"

	"github.com/Siravitt/go-todoist/errors"
	"github.com/Siravitt/go-todoist/logs"
	"github.com/Siravitt/go-todoist/repository/user_repository"
	"github.com/Siravitt/go-todoist/utils"
	"github.com/go-playground/validator/v10"
)

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) authService {
	return authService{userRepo: userRepo}
}

func (s authService) Login(login LoginRequest) (*string, error) {
	user, err := s.userRepo.GetByEmail(login.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User not found")
		}
		return nil, err
	}
	// Logic for compare password
	if !utils.IsValidPassword(user.Password, login.Password) {
		return nil, errors.NewUnauthorizedError("Invalid email or password")
	}
	// Logic for compare password
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s authService) Register(regis RegisterRequest) error {
	validate := validator.New()
	validate.RegisterValidation("phoneNumber", utils.IsValidPhone)

	err := validate.Struct(regis)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			errorMessage := "Field:" + e.Field() + "failed validation:" + e.Tag()
			logs.Error(errorMessage)
		}
		return errors.NewValidationError("Register information is required")
	}
	regis.Password, err = utils.HashPassword(regis.Password)
	if err != nil {
		logs.Error(err)
		return errors.NewUnexpectedError()
	}
	registerStruct := user_repository.User{
		Id:          0,
		Username:    regis.Username,
		Password:    regis.Password,
		Email:       regis.Email,
		PhoneNumber: regis.PhoneNumber,
	}
	_, err = s.userRepo.Create(registerStruct)
	if err != nil {
		logs.Error(err)
		return errors.NewUnexpectedError()
	}
	return nil
}
