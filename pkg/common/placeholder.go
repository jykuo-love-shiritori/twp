package common

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var DefaultImageUuid = pgtype.UUID{
	Bytes: [16]byte{},
	Valid: true,
}
