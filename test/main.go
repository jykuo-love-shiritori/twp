package main

import (
	"context"
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
)

func main() {
	var err error
	pg, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	err = pg.Queries.InsertTestUser(context.Background(), pgtype.UUID{Valid: true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InsertTestUser success")
	err = pg.Queries.DeleteTestUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeleteTestUser success")

	AddMockUsers(pg)

}

func AddMockUsers(pg *db.DB) {
	for i := 0; i < 10; i++ {
		startRune, _ := utf8.DecodeRuneInString("🐱")
		avatar := string(startRune + rune(i))
		mockData := db.AddUserParams{
			Username: fmt.Sprintf("test%d", i),
			Password: fmt.Sprintf("test%d", i),
			Name:     avatar,
			Email:    fmt.Sprintf("test%d", i) + "@test.com",
			ImageID:  pgtype.UUID{Valid: true},
		}
		if err := pg.Queries.AddUser(context.Background(), mockData); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertMockUser success")
}
