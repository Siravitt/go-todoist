package todo_repository

import "github.com/jmoiron/sqlx"

type todoRepositoryDB struct {
	db *sqlx.DB
}

func NewTodoRepositoryDB(db *sqlx.DB) TodoRepository {
	return todoRepositoryDB{db: db}
}

// Get all todos by user_id
func (r todoRepositoryDB) GetAll(userID int) ([]Todo, error) {
	todos := []Todo{}
	query := "select id, title, description, user_id from todos where user_id = ?"
	err := r.db.Select(&todos, query, userID)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Get todo by todo_id
func (r todoRepositoryDB) GetByID(id int) (*Todo, error) {
	todo := Todo{}
	query := "select id, title, description, user_id from todos where id = ?"
	err := r.db.Get(&todo, query, id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Create todo
func (r todoRepositoryDB) Create(t Todo) (*Todo, error) {
	query := "insert into todos (title, description, completed, user_id) values (?, ?, ?, ?)"
	result, err := r.db.Exec(query, t.Title, t.Description, t.Completed, t.UserID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.Id = int(id)

	return &t, nil
}

// Update todo
func (r todoRepositoryDB) Update(t Todo) (*Todo, error) {
	return nil, nil
}

// Delete todo by id
func (r todoRepositoryDB) Delete(id int) error {
	return nil
}
