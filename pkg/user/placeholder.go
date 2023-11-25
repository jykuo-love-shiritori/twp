package user

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var defaultImageUuid = pgtype.UUID{
	Bytes: [16]byte{},
	Valid: true,
}
