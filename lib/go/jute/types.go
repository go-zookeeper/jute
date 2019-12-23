package jute

import "errors"

var ErrNilKey = errors.New("got nil value for map key")

// Returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}
