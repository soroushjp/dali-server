package context

// AppContext holds shared context for the server.
import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Imported for side effects
	"github.com/plaid/go-envvar/envvar"
)

// AppContext provides the context for the server. A database connection and
// parsed environment variables are provided.
type AppContext struct {
	DB  *sqlx.DB
	Env serverEnvVars
}

type serverEnvVars struct {
	Port   uint   `envvar:"DALI_PORT"`
	DBUser string `envvar:"DALI_DB_USER"`
	DBName string `envvar:"DALI_DB_NAME_PREFIX"`
}

// NewAppContext initializes and returns a new app context.
func NewAppContext() (*AppContext, error) {
	// Get enviornment variables
	env := serverEnvVars{}
	if err := envvar.Parse(&env); err != nil {
		return nil, err
	}
	// Get DB connection
	db, err := sqlx.Connect(
		"postgres",
		"user="+env.DBUser+
			" dbname="+env.DBName+
			" sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	return &AppContext{
		DB:  db,
		Env: env,
	}, nil
}
