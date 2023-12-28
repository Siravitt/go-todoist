package user_service

type UserResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
}
