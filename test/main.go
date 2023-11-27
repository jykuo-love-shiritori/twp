package main

import (
	"context"
	"encoding/json"
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
		emptyJSON := map[string]interface{}{
			"card_number": fmt.Sprint(i),
			"exp_month":   fmt.Sprint(i),
			"exp_year":    fmt.Sprint(i),
			"cvc":         fmt.Sprint(i),
		}
		// Marshal the empty JSON object into a JSON-formatted byte slice
		jsonData, err := json.Marshal(emptyJSON)
		if err != nil {
			panic(err)
		}
		startRune, _ := utf8.DecodeRuneInString("🐱")
		avatar := string(startRune + rune(i))
		mockData := db.AddUserParams{
			SellerName: fmt.Sprintf("test%d", i),
			Password:   fmt.Sprintf("test%d", i),
			Name:       avatar,
			Email:      fmt.Sprintf("test%d", i) + "@test.com",
			Address:    fmt.Sprintf("test%d", i),
			ImageID:    pgtype.UUID{Valid: true},
			Role:       "customer",
			CreditCard: jsonData,
		}
		err = pg.Queries.AddUser(context.Background(), mockData)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("InsertMockUser success")
}
