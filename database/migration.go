package database

import (
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

const MigrationSource string = "file://migrations"

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
