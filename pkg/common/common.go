package common

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

type QueryParams struct {
	Offset int32 `query:"offset"` // the start index of the query
	Limit  int32 `query:"limit"`  // the number of the query
}

type Cart struct {
	Seller_name string
	Products    []db.GetProductInCartRow
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
	newFileName := id.String() + filepath.Ext(file.Filename)
	return newFileName
}

func FileMimeFrom(fileName string) string {
	switch filepath.Ext(fileName) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".svg":
		return "image/svg+xml"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream" // default content type
	}
}
