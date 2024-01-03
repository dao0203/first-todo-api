package repository

import (
	entity "first-todo-api/domain"
	postgres "first-todo-api/postgres"
	"sync"
)

type todoRepositoryImpl struct {
	postgresHandler postgres.PostgresQuery
}

var (
	todoRepositoryInstance *todoRepositoryImpl
	todoRepositoryOnce     sync.Once
)

func NewTodoRepository(handler postgres.PostgresQuery) TodoRepository {
	todoRepositoryOnce.Do(func() {
		todoRepositoryInstance = &todoRepositoryImpl{postgresHandler: handler}
	})
	return todoRepositoryInstance
}

func (repo *todoRepositoryImpl) Delete(todo entity.Todo) error {
	err := repo.postgresHandler.DeleteTodo(todo)

	return err
}

func (repo *todoRepositoryImpl) Update(todo entity.Todo) error {
	err := repo.postgresHandler.UpdateTodo(todo)

	return err
}

func (*todoRepositoryImpl) Create(todo entity.Todo) error {
	err := todoRepositoryInstance.postgresHandler.CreateTodo(todo)

	return err
}

func (*todoRepositoryImpl) FindAll() ([]entity.Todo, error) {
	todos, err := todoRepositoryInstance.postgresHandler.FindAllTodo()

	return todos, err
}

func (*todoRepositoryImpl) FindByID(id int) (entity.Todo, error) {
	todo, err := todoRepositoryInstance.postgresHandler.FindTodoByID(id)

	return todo, err
}
