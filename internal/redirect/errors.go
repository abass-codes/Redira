package redirect

import "errors"

var (
	ErrNotFound = errors.New("link not found")
	ErrDisabled = errors.New("link disabled")
	ErrExpired  = errors.New("link expired")
)
