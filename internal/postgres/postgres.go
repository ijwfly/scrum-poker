package postgres

import (
	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{db: db}
}
