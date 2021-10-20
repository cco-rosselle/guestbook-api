package repos

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type PostgresRepo struct {
	log zerolog.Logger
	pool *sql.DB
}

func newPostgresRepo() (*PostgresRepo, error) {
	//s := settings.Default()

	connection := "user=postgres dbname=guestbook-db password=123 host=localhost sslmode=disable"

	_, cancel := getContext()
	defer cancel()

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Error().
			Stack().
			Err(err).
			Msg("unable to connect to postgres cluster")

		return nil, err
	}

	log.Trace().Msg("connected to postgres cluster")

	return &PostgresRepo {
		log: log.With().
			Str("package", "repos").
			Logger(),
		pool: db,
	}, nil
} 