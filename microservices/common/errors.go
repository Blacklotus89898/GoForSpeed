package common

import "errors"


var (
	ErrNotFound = errors.New("not found")
	ErrNoItems = errors.New("no items in order")
)

