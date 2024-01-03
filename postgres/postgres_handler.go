package infra

import (
	"database/sql"
	entity "first-todo-api/domain"

	_ "github.com/lib/pq"
)

type PostgresHandler interface {
	Create(todo entity.Todo) (err error)
	FindAll() (todos []entity.Todo, err error)
	FindByID(id int) (todo entity.Todo, err error)
	Update(todo entity.Todo) (err error)
	Delete(todo entity.Todo) (err error)
}

type postgresHandlerImpl struct {
	Conn *sql.DB
}

func NewPostgresHandler(conn *sql.DB) PostgresHandler {
	return &postgresHandlerImpl{Conn: conn}
}

func (handler *postgresHandlerImpl) Create(todo entity.Todo) (err error) {
	//titleだけを使ってINSERT文を実行している
	_, err = handler.Conn.Exec("INSERT INTO todos (title) VALUES ($1)", todo.Title)
	return
}

func (handler *postgresHandlerImpl) FindAll() (todos []entity.Todo, err error) {
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

func (handler *postgresHandlerImpl) FindByID(id int) (todo entity.Todo, err error) {
	rows, err := handler.Conn.Query("SELECT * FROM todos WHERE id = $1", id)
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

func (handler *postgresHandlerImpl) Update(todo entity.Todo) (err error) {
	_, err = handler.Conn.Exec("UPDATE todos SET title = $1, done = $2 WHERE id = $3", todo.Title, todo.Done, todo.ID)
	return
}

func (handler *postgresHandlerImpl) Delete(todo entity.Todo) (err error) {
	_, err = handler.Conn.Exec("DELETE FROM todos WHERE id = $1", todo.ID)
	return
}
