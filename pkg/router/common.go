package router

import (
	"regexp"

	"github.com/jackc/pgx/v5"
)

type Failure struct {
	Msg string `json:"message"`
}

func hasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func PGXErrorExtraction(err error) string {
	switch err {
	case nil:
		return "success"
	case pgx.ErrNoRows:
		return "Not Found"
	case pgx.ErrTooManyRows:
		return "Too Many Result"
	default:
		return "Internal Server Error"
	}
}
