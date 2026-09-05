package implementation

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrTTLExpired  = errors.New("ttl expired")
	ErrInvalidTTL  = errors.New("ttl must be greater than zero")
)

