package repository

type Repository struct {
	Authorisation
	TodoItem
	TodoList
}

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

func NewRepository() *Repository {
	return &Repository{}
}
