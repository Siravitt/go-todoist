package user_repository

type User struct {
	Id          int    `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	Email       string `db:"email"`
	PhoneNumber int    `db:"phone_number"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(int) (*User, error)
	Create(User) (*User, error)
}
