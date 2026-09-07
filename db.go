package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store defines the persistence operations a paste storage backend must support.
type Store interface {
	// Save persists a new paste.
	Save(paste Paste) error
	// Get retrieves a paste by its ID.
	Get(id string) (Paste, error)
}

// PostgresStore is a Store implementation backed by PostgreSQL.
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore opens a connection to Postgres using dsn and verifies
// it is reachable before returning a ready-to-use PostgresStore.
func NewPostgresStore(dsn string) (*PostgresStore, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: conn}, nil
}
