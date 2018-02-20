package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/HackGT/SponsorshipPortal/config"
)

func New(config *config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
