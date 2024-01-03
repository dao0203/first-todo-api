package infra

import (
	"database/sql"
	entity "first-todo-api/domain"

	_ "github.com/lib/pq"
)

type PostgresQuery interface {
	CreateTodo(todo entity.Todo) (err error)
	FindAllTodo() (todos []entity.Todo, err error)
	FindTodoByID(id int) (todo entity.Todo, err error)
	UpdateTodo(todo entity.Todo) (err error)
	DeleteTodo(todo entity.Todo) (err error)
}

type postgresQueryImpl struct {
	Conn *sql.DB
}

func NewPostgresQuery(conn *sql.DB) PostgresQuery {
	return &postgresQueryImpl{Conn: conn}
}

func (query *postgresQueryImpl) CreateTodo(todo entity.Todo) (err error) {
	_, err = query.Conn.Exec("INSERT INTO todos (title) VALUES ($1)", todo.Title)
	return
}

func (handler *postgresQueryImpl) FindAllTodo() (todos []entity.Todo, err error) {
	rows, err := handler.Conn.Query("SELECT * FROM todos")
	if err != nil {
		return
	}
	for rows.Next() {
		var todo entity.Todo
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			return
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

func (query *postgresQueryImpl) FindTodoByID(id int) (todo entity.Todo, err error) {
	rows, err := query.Conn.Query("SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			return
		}
	}
	return
}

func (query *postgresQueryImpl) UpdateTodo(todo entity.Todo) (err error) {
	_, err = query.Conn.Exec("UPDATE todos SET title = $1, done = $2 WHERE id = $3", todo.Title, todo.Done, todo.ID)
	return
}

func (query *postgresQueryImpl) DeleteTodo(todo entity.Todo) (err error) {
	_, err = query.Conn.Exec("DELETE FROM todos WHERE id = $1", todo.ID)
	return
}
