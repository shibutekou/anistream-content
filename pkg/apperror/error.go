package apperror

import "errors"

var (
	ErrTitleNotFound         = errors.New("title with received parameters not found")
	ErrMissingOrInvalidToken = errors.New("missing or invalid token")
	ErrNoSearchParams        = errors.New("at least one search parameter was not specified")
	ErrUnknown               = errors.New("unknown error")
)
