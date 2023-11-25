package common

// This is used as a placeholder for `echo.HTTPError`
// because echo-swagger cannot find external types
type HttpError struct {
	Message interface{} `json:"message"`
}
