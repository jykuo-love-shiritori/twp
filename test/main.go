package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

	testAddUser(pg)

}

func testAddUser(pg *db.DB) {
	for i := 0; i < 100; i++ {
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
		mockData := db.AddUserParams{
			Username:   fmt.Sprintf("test%d", i),
			Password:   fmt.Sprintf("test%d", i),
			Name:       fmt.Sprintf("test%d", i),
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
	fmt.Println("InsertUser success")
}
