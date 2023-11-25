package db

type JSONB map[string]interface{}
type creditCard struct {
	Name       string `json:"name"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"CVV"`
}
