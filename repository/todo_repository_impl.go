package repository

import (
	entity "first-todo-api/internal/domain"
	"sync"
)

type todoRepositoryImpl struct{}

var (
	todoRepositoryInstance *todoRepositoryImpl
	todoRepositoryOnce     sync.Once
)

func NewTodoRepository() TodoRepository {
	todoRepositoryOnce.Do(func() {
		todoRepositoryInstance = &todoRepositoryImpl{}
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
