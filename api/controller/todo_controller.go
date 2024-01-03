package controller

import (
	"first-todo-api/entity"
	usecase "first-todo-api/usecase"
)

type TodoController struct {
	TodoUsecase usecase.TodoUsecase
}

func NewTodoController(todoUsecase usecase.TodoUsecase) *TodoController {
	return &TodoController{
		TodoUsecase: todoUsecase,
	}
}

func (tc *TodoController) Create() {
	var task entity.Todo

}
