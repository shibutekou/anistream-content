package apperror

import "errors"

var (
	ErrTitleNotFound = errors.New("title with received parameters not found")
)
