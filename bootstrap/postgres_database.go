package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresDatabase struct {
	conn *sql.DB
}

func NewPostgresDatabase(env *Env) *PostgresDatabase {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		env.PostgresHost,
		env.PostgresPort,
		env.PostgresUser,
		env.PostgresDBName,
		env.PostgresPassword,
	)
	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &PostgresDatabase{
		conn: conn,
	}
}

func ClosePostgresDatabase(db *PostgresDatabase) {
	db.conn.Close()
	log.Println("Postgres Database Closed")
}
