package bootstrap

type Applicaton struct {
	Env              *Env
	PostgresDatabase *PostgresDatabase
}

func NewApplication() *Applicaton {
	env := NewEnv()
	postgresDatabase := NewPostgresDatabase(env)
	return &Applicaton{
		Env:              env,
		PostgresDatabase: postgresDatabase,
	}
}
