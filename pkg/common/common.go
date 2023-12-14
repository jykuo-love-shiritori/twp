package common

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

type QueryParams struct {
	Offset int32 `query:"offset"` // the start index of the query
	Limit  int32 `query:"limit"`  // the number of the query
}

func NewQueryParams(offset int32, limit int32) QueryParams {
	return QueryParams{Offset: offset, Limit: limit}
}

func (q *QueryParams) Validate() error {
	if q.Offset < 0 || q.Limit < 0 || q.Limit > constants.QUERY_LIMIT {
		return errors.New("invalid query parameter")
	}
	return nil
}
func HasRegexSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

// c.FormValue("tags")
func String2IntArray(str string) ([]int32, error) {
	var array []int32
	err := json.Unmarshal([]byte(str), &array)
	if err != nil {
		return nil, err
	}
	return array, nil
}
func GetFileName(file *multipart.FileHeader) string {
	id := uuid.New()
	parts := strings.Split(file.Filename, ".")
	fileType := parts[len(parts)-1]
	newFileName := id.String() + "." + fileType
	return newFileName
}
