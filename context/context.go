package context

// AppContext holds shared context for the server.
import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AppContext struct {
	DB *sqlx.DB
}

// NewAppContext initializes and returns a new app context.
func NewAppContext() (*AppContext, error) {
	db, err := sqlx.Connect("postgres", "user=dali dbname=dali sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &AppContext{
		DB: db,
	}, nil
}
