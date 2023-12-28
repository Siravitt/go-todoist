package user_repository

import "github.com/jmoiron/sqlx"

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepositoryDB{db: db}
}

// Get all users
func (r userRepositoryDB) GetAll() ([]User, error) {
	users := []User{}
	query := "select id, username, password, email, phone_number from users"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Get user by user_id
func (r userRepositoryDB) GetByID(id int) (*User, error) {
	user := User{}
	query := "select id, username, password, email, phone_number from users where id = ?"
	err := r.db.Get(&user, query)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create user
func (r userRepositoryDB) Create(u User) (*User, error) {
	query := "insert into users (username, password, email, phone_number) values (?, ?, ?, ?)"
	result, err := r.db.Exec(query, u.Username, u.Password, u.Email, u.PhoneNumber)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = int(id)

	return &u, nil
}
