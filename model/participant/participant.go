package participant

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/HackGT/SponsorshipPortal/model"
)

const (
	table = "participants"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

type Participant struct {
	ID             int64  `db:"id"`
	RegistrationID string `db:"registration_id"`
}

func ByID(db model.Connection, id int64) (Participant, bool, error) {
	result := Participant{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, registration_id
FROM %v
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), id)
	return result, err != sql.ErrNoRows, err
}

func ByRegistrationID(db model.Connection, regID string) (Participant, bool, error) {
	result := Participant{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, registration_id
FROM %v
WHERE registration_id = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), regID)
	return result, err != sql.ErrNoRows, err
}

func MatchKeywords(db model.Connection, keywords ...string) ([]Participant, error) {
	return nil, ErrNotImplemented
}

func Create(db model.Connection, regID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
INSERT INTO %v
(registration_id)
VALUES
($1)`,
		table), regID)
	return result, err
}

func (p *Participant) SetDocumentContents(db model.Connection, doc string) (sql.Result, error) {
	return nil, ErrNotImplemented
}
