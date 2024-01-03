package router

import (
	"first-todo-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.PostgresDatabase, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	NewTodoRouter(env, timeout, db, publicRouter)
}
