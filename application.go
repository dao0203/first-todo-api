package main

import (
	"first-todo-api/api/router"
	"first-todo-api/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.NewApplication()

	db := app.PostgresDatabase
	defer db.ClosePostgresDatabase()

	timeout := time.Duration(5) * time.Second

	gin := gin.Default()

	router.Setup(timeout, db, gin)

	gin.Run()
}
