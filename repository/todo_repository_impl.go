package repository

import (
	entity "first-todo-api/domain"
	postgres "first-todo-api/postgres"
	"sync"
)

type todoRepositoryImpl struct {
	postgresHandler postgres.PostgresHandler
}

var (
	todoRepositoryInstance *todoRepositoryImpl
	todoRepositoryOnce     sync.Once
)

func NewTodoRepository(handler postgres.PostgresHandler) TodoRepository {
	todoRepositoryOnce.Do(func() {
		todoRepositoryInstance = &todoRepositoryImpl{postgresHandler: handler}
	})
	return todoRepositoryInstance
}

func (*todoRepositoryImpl) Delete(todo entity.Todo) error {
	panic("unimplemented")
}

func (*todoRepositoryImpl) Update(todo entity.Todo) error {
	panic("unimplemented")
}

func (*todoRepositoryImpl) Create(todo entity.Todo) error {
	panic("unimplemented")
}

func (*todoRepositoryImpl) FindAll() ([]entity.Todo, error) {
	panic("unimplemented")
}

func (*todoRepositoryImpl) FindByID(id int) (entity.Todo, error) {
	panic("unimplemented")
}
