package router

import (
	"regexp"
)

type searchParams struct {
	Limit  int32 `json:"limit" query:"limit"`
	Offset int32 `json:"offset" query:"offset"`
}

func hasRegexSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}
