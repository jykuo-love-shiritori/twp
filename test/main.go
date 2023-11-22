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
	for i := 0; i < 100; i++ {
		emptyJSON := map[string]interface{}{
			"test": "test",
		}

		// Marshal the empty JSON object into a JSON-formatted byte slice
		jsonData, err := json.Marshal(emptyJSON)
		if err != nil {
			panic(err)
		}
		fmt.Println((string(jsonData)))
		mockData := db.AddUserParams{
			Username:   "test" + string(i),
			Password:   "test" + string(i),
			Name:       "test" + string(i),
			Email:      "test" + string(i) + "@test.com",
			Address:    "test" + string(i),
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
