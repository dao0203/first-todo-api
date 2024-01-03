package api

import (
	"first-todo-api/api/controller"
	"first-todo-api/bootstrap"
	"first-todo-api/postgres"
	"first-todo-api/repository"
	"first-todo-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewTodoRouter(env *bootstrap.Env, timeout time.Duration, grou *gin.RouterGroup) {
	database := bootstrap.NewPostgresDatabase(env)
	query := postgres.NewPostgresQuery(database.Conn)
	todoRepository := repository.NewTodoRepository(query)
	todoUsecase := usecase.NewTodoUsecase(todoRepository, timeout)
	todoController := &controller.TodoController{TodoUseCase: todoUsecase}

	grou.POST("/todos", todoController.Create)
	grou.GET("/todos", todoController.FindAll)
	grou.GET("/todos/:id", todoController.FindByID)
	grou.PUT("/todos/:id", todoController.Update)
	grou.DELETE("/todos/:id", todoController.Delete)
}
