package usecase

import (
	"context"
	entity "first-todo-api/domain"
	"first-todo-api/repository"
	"time"
)

type TodoUsecase interface {
	Create(c context.Context, todo entity.Todo) (err error)
	FindAll(c context.Context) (todos []entity.Todo, err error)
	FindByID(c context.Context, id int) (todo entity.Todo, err error)
	Update(c context.Context, todo entity.Todo) (err error)
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

func (us *todoUsecase) Create(c context.Context, todo entity.Todo) (err error) {
	_, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	err = us.todoRepository.Create(todo)

	return err
}

func (us *todoUsecase) FindAll(c context.Context) (todos []entity.Todo, err error) {
	_, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	todos, err = us.todoRepository.FindAll()

	return todos, err
}

func (us *todoUsecase) FindByID(c context.Context, id int) (todo entity.Todo, err error) {
	_, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	todo, err = us.todoRepository.FindByID(id)

	return todo, err
}

func (us *todoUsecase) Update(c context.Context, todo entity.Todo) (err error) {
	_, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = us.todoRepository.Update(todo)

	return err
}
