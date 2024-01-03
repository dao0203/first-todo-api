package usecase

import (
	entity "first-todo-api/domain"
	"first-todo-api/repository"
	"time"
)

type TodoUsecase interface {
	Create(todo entity.Todo) (err error)
	FindAll() (todos []entity.Todo, err error)
	FindByID(id int) (todo entity.Todo, err error)
	Update(todo entity.Todo) (err error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
	contextTimeout time.Duration
}

func NewTodoUsecase(todoRepository repository.TodoRepository, contextTimeout time.Duration) TodoUsecase {
	return &todoUsecase{
		todoRepository: todoRepository,
		contextTimeout: contextTimeout,
	}
}

func (*todoUsecase) Create(todo entity.Todo) (err error) {
	panic("unimplemented")
}

func (*todoUsecase) FindAll() (todos []entity.Todo, err error) {
	panic("unimplemented")
}

func (*todoUsecase) FindByID(id int) (todo entity.Todo, err error) {
	panic("unimplemented")
}

func (*todoUsecase) Update(todo entity.Todo) (err error) {
	panic("unimplemented")
}
