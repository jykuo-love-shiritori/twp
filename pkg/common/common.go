package common

import (
	"regexp"

	"github.com/jykuo-love-shiritori/twp/db"
)

type Failure struct {
	Msg string `json:"message"`
}

type QueryParams struct {
	Offset int32 `query:"offset"` // the start index of the query
	Limit  int32 `query:"limit"`  // the number of the query
}

func NewQueryParams(offset int32, limit int32) QueryParams {
	return QueryParams{Offset: offset, Limit: limit}
}

type Cart struct {
	Seller_name string
	Products    []db.GetProductInCartRow
}

func HasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}
