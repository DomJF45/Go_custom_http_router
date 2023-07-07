package routererrors

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("no matching route was found")
	ErrMethodNotAllowed = errors.New("method is not allowed")
)
