package database

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	"github.com/HackGT/SponsorshipPortal/config"
)

const MigrationSource string = "file://migrations"

func loadFromConfig(db *sql.DB, config *config.DatabaseConfig) (*migrate.Migrate, error) {
	pgDriver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: postgres.DefaultMigrationsTable,
		DatabaseName:    config.DbName,
	})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		MigrationSource,
		"postgres", pgDriver,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Migrate(connString string) error {
	m, err := migrate.New(
		MigrationSource,
		connString,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	_, err = m.Close()
	return err
}

func IsReadyWithInstance(db *sql.DB, config *config.DatabaseConfig) (bool, error) {
	m, err := loadFromConfig(db, config)
	if err != nil {
		if err == postgres.ErrDatabaseDirty {
			return false, nil
		} else {
			return false, err
		}
	}

	_, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return false, nil
		} else {
			return false, err
		}
	}
	return !dirty, nil
}
