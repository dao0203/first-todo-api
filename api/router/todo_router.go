package router

import (
	"first-todo-api/api/controller"
	"first-todo-api/bootstrap"
	"first-todo-api/postgres"
	"first-todo-api/repository"
	"first-todo-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewTodoRouter(env *bootstrap.Env, timeout time.Duration, database *bootstrap.PostgresDatabase, group *gin.RouterGroup) {
	query := postgres.NewPostgresQuery(database.Conn)
	todoRepository := repository.NewTodoRepository(query)
	todoUsecase := usecase.NewTodoUsecase(todoRepository, timeout)
	todoController := &controller.TodoController{TodoUseCase: todoUsecase}

	group.POST("/todos", todoController.Create)
	group.GET("/todos", todoController.FindAll)
	group.GET("/todos/:id", todoController.FindByID)
	group.PUT("/todos/:id", todoController.Update)
	group.DELETE("/todos/:id", todoController.Delete)
}
