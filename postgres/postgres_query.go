package postgres

import (
	"database/sql"
	"first-todo-api/entity"

	_ "github.com/lib/pq"
)

type PostgresQuery interface {
	CreateTodo(todo entity.Todo) (err error)
	FindAllTodo() (todos []entity.Todo, err error)
	FindTodoByID(id string) (todo entity.Todo, err error)
	UpdateTodo(todo entity.Todo) (err error)
	DeleteTodo(todo entity.Todo) (err error)
}

type postgresQueryImpl struct {
	conn *sql.DB
}

func NewPostgresQuery(conn *sql.DB) PostgresQuery {

	return &postgresQueryImpl{conn: conn}
}

func (query *postgresQueryImpl) CreateTodo(todo entity.Todo) (err error) {
	_, err = query.conn.Exec("INSERT INTO todos (title) VALUES ($1)", todo.Title)
	return
}

func (handler *postgresQueryImpl) FindAllTodo() (todos []entity.Todo, err error) {
	rows, err := handler.conn.Query("SELECT * FROM todos")
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

func (query *postgresQueryImpl) FindTodoByID(id string) (todo entity.Todo, err error) {
	rows, err := query.conn.Query("SELECT * FROM todos WHERE id = $1", id)
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
	_, err = query.conn.Exec("UPDATE todos SET title = $1, done = $2 WHERE id = $3", todo.Title, todo.Done, todo.ID)
	return
}

func (query *postgresQueryImpl) DeleteTodo(todo entity.Todo) (err error) {
	_, err = query.conn.Exec("DELETE FROM todos WHERE id = $1", todo.ID)
	return
}
