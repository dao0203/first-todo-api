package controller

import (
	"first-todo-api/entity"
	"first-todo-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoUseCase usecase.TodoUsecase
}

func NewTodoController(todoUsecase usecase.TodoUsecase) *TodoController {
	return &TodoController{
		TodoUseCase: todoUsecase,
	}
}

func (tc *TodoController) Create(c *gin.Context) {
	var task entity.Todo

	err := c.ShouldBind(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TodoUseCase.Create(c, task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)

	c.JSON(http.StatusOK, entity.SuccessResponse{Message: "Todo created successfully"})
}

func (tc *TodoController) FindAll(c *gin.Context) {
	todos, err := tc.TodoUseCase.FindAll(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) FindByID(c *gin.Context) {

	id := c.Param("id")

	todo, err := tc.TodoUseCase.FindByID(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) Update(c *gin.Context) {
	var task entity.Todo

	err := c.ShouldBind(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TodoUseCase.Update(c, task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.SuccessResponse{Message: "Todo updated successfully"})
}

func (tc *TodoController) Delete(c *gin.Context) {

	id := c.Param("id")

	todo := entity.Todo{ID: id}

	err := tc.TodoUseCase.Delete(c, todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.SuccessResponse{Message: "Todo deleted successfully"})
}
