package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID: "twp",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8080/api/oauth/authorize",
			TokenURL: "http://localhost:8080/api/oauth/token",
		},
	}

	verifier := oauth2.GenerateVerifier()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	fmt.Print("\nEnter code: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	resp, err := client.Get("http://localhost:8080/api/ping")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Println(string(body))
}
