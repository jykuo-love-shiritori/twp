package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	q    *Queries
	Conn *pgx.Conn
}

func NewDB() (*DB, error) {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	q := New(conn)

	return &DB{q: q, Conn: conn}, nil
}

func (db *DB) Close() {
	db.Conn.Close(context.Background())
}
