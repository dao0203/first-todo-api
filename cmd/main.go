package main

import (
	"first-todo-api/api/router"
	"first-todo-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.NewApplication()
	env := app.Env

	db := app.PostgresDatabase
	defer db.ClosePostgresDatabase()

	timeout := time.Duration(5) * time.Second

	gin := gin.Default()

	router.Setup(env, timeout, db, gin)

	gin.Run("http://dao0203-first-todo-api.com") //仮のドメイン
}
