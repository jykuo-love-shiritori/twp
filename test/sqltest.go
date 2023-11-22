package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
)

func insertTest() {
	var err error
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Queries.InsertTestUser(context.Background(), pgtype.UUID{Valid: true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InsertTestUser success")
	err = db.Queries.DeleteTestUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeleteTestUser success")
}
