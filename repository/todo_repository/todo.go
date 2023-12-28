package todo_repository

type Todo struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Completed   int    `db:"completed"`
	UserID      int    `db:"user_id"`
}

type TodoRepository interface {
	GetAll(int) ([]Todo, error)
	GetByID(int) (*Todo, error)
	Create(Todo) (*Todo, error)
	Update(Todo) (*Todo, error)
	Delete(int) error
}
