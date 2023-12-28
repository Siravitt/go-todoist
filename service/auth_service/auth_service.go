package auth_service

import "github.com/Siravitt/go-todoist/repository/user_repository"

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) authService {
	return authService{userRepo: userRepo}
}

func (s authService) Login(login LoginRequest) {

}

func (s authService) Register(regis RegisterRequest) {

}
