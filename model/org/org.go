package org

import (
	"database/sql"
	"fmt"

	"github.com/HackGT/SponsorshipPortal/model"
)

const (
	table = "sponsor_orgs"
)

type Org struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func ByID(db model.Connection, id int64) (Org, bool, error) {
	result := Org{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, name
FROM %v
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), id)
	return result, err != sql.ErrNoRows, err
}

func ByName(db model.Connection, name string) (Org, bool, error) {
	result := Org{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, name
FROM %v
WHERE name = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), name)
	return result, err != sql.ErrNoRows, err
}

func Create(db model.Connection, name string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
INSERT INTO %v
(name)
VALUES
($1)`,
		table), name)
	return result, err
}
