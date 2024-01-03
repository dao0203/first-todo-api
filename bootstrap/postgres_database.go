package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresDatabase struct {
	Conn *sql.DB
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
		Conn: conn,
	}
}

func (db *PostgresDatabase) ClosePostgresDatabase() {
	db.Conn.Close()
	log.Println("Postgres Database Closed")
}
