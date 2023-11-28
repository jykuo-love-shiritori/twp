// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: test_temp.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteTestUser = `-- name: DeleteTestUser :exec

DELETE FROM "user" WHERE "username" = ' user0 '
`

func (q *Queries) DeleteTestUser(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteTestUser)
	return err
}

const insertTestUser = `-- name: InsertTestUser :exec

INSERT INTO
    "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card"
    )
VALUES (
        'user0',
        'password0',
        'name0',
        'email0',
        'address0',
        $1,
        'customer',
        '{"card_number": "card_number0", "expiration_date": "expiration_date0", "cvv": "cvv0"}'
    )
`

func (q *Queries) InsertTestUser(ctx context.Context, imageID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, insertTestUser, imageID)
	return err
}

const searchTestUser = `-- name: SearchTestUser :one

SELECT id, username, password, name, email, address, image_id, role, credit_card, enabled FROM "user" WHERE "username" = ' user0 '
`

func (q *Queries) SearchTestUser(ctx context.Context) (User, error) {
	row := q.db.QueryRow(ctx, searchTestUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageID,
		&i.Role,
		&i.CreditCard,
		&i.Enabled,
	)
	return i, err
}
