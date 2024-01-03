package repository

import (
	"first-todo-api/entity"
	postgres "first-todo-api/postgres"
	"sync"
)

type todoRepositoryImpl struct {
	postgresQuery postgres.PostgresQuery
}

var (
	todoRepositoryInstance *todoRepositoryImpl
	todoRepositoryOnce     sync.Once
)

func NewTodoRepository(query postgres.PostgresQuery) TodoRepository {
	todoRepositoryOnce.Do(func() {
		todoRepositoryInstance = &todoRepositoryImpl{postgresQuery: query}
	})
	return todoRepositoryInstance
}

func (repo *todoRepositoryImpl) Delete(todo entity.Todo) error {
	err := repo.postgresQuery.DeleteTodo(todo)

	return err
}

func (repo *todoRepositoryImpl) Update(todo entity.Todo) error {
	err := repo.postgresQuery.UpdateTodo(todo)

	return err
}

func (*todoRepositoryImpl) Create(todo entity.Todo) error {
	err := todoRepositoryInstance.postgresQuery.CreateTodo(todo)

	return err
}

func (*todoRepositoryImpl) FindAll() ([]entity.Todo, error) {
	todos, err := todoRepositoryInstance.postgresQuery.FindAllTodo()

	return todos, err
}

func (*todoRepositoryImpl) FindByID(id string) (entity.Todo, error) {
	todo, err := todoRepositoryInstance.postgresQuery.FindTodoByID(id)

	return todo, err
}
