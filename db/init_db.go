package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Queries *Queries
	Conn    *pgx.Conn
}

func NewDB() (*DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbname)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	q := New(conn)

	return &DB{Queries: q, Conn: conn}, nil
}

func (db *DB) Close() {
	err := db.Conn.Close(context.Background())
	if err != nil {
		panic(err)
	}
}
