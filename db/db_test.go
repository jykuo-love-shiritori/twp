package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestInsert(t *testing.T) {
	var err error
	db, _ := NewDB()
	defer db.Close()

	err = db.q.InsertTestUser(context.Background(), pgtype.UUID{Valid: true})
	if err != nil {
		t.Fatal(err)
	}

	err = db.q.DeleteTestUser(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
