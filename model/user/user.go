package user

import (
	"database/sql"
	"fmt"

	"github.com/HackGT/SponsorshipPortal/model"
)

const (
	table = "sponsors"
)

type User struct {
	ID       int64  `db:"id"`
	OrgID    int64  `db:"org_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	State    string `db:"state"`
}

func ByEmail(db model.Connection, email string) (User, bool, error) {
	result := User{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, org_id, email, password, state
FROM %v
WHERE email = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), email)
	return result, err != sql.ErrNoRows, err
}

func ByID(db model.Connection, id int64) (User, bool, error) {
	result := User{}
	err := db.Get(&result, fmt.Sprintf(`
SELECT id, org_id, email, password, state
FROM %v
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1`,
		table), id)
	return result, err != sql.ErrNoRows, err
}

func Create(db model.Connection, orgID int64, email, password string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
INSERT INTO %v
(org_id, email, password)
VALUES
($1, $2, $3)`,
		table), orgID, email, password)
	return result, err
}

func (u *User) Save(db model.Connection) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
UPDATE %v
   SET org_id = $1,
       email = $2,
	   password = $3,
	   state = $4
 WHERE id = $5;
`,
		table), u.OrgID, u.Email, u.Password, u.State, u.ID)
	return result, err
}
