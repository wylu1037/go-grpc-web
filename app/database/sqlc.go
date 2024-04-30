package database

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/rs/zerolog/log"
)

//go:embed sql/schema.sql
var ddl string

func NewQueries() *Queries {
	// open db file
	db, err := sql.Open("sqlite3", "file:./data.db")
	if err != nil {
		log.Error().Err(err).Msg("failed to open database")
	}

	// create tables
	if _, err := db.ExecContext(context.Background(), ddl); err != nil {
		log.Error().Err(err).Msg("failed to execute ddl")
	}

	return New(db)
}
