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

func (repo *todoRepositoryImpl) Delete(todo entity.Todo) error {
	err := repo.postgresHandler.Delete(todo)

	if err != nil {
		return err
	}

	return nil
}

func (repo *todoRepositoryImpl) Update(todo entity.Todo) error {
	err := repo.postgresHandler.Update(todo)
	if err != nil {
		return err
	}

	return nil
}

func (*todoRepositoryImpl) Create(todo entity.Todo) error {
	err := todoRepositoryInstance.postgresHandler.Create(todo)
	if err != nil {
		return err
	}

	return nil
}

func (*todoRepositoryImpl) FindAll() ([]entity.Todo, error) {
	todos, err := todoRepositoryInstance.postgresHandler.FindAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (*todoRepositoryImpl) FindByID(id int) (entity.Todo, error) {
	todo, err := todoRepositoryInstance.postgresHandler.FindByID(id)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}
