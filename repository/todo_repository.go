package repository

import entity "first-todo-api/internal/domain"

type TodoRepository interface {
	Create(todo entity.Todo) error
	FindAll() ([]entity.Todo, error)
	FindByID(id int) (entity.Todo, error)
	Update(todo entity.Todo) error
	Delete(todo entity.Todo) error
}
