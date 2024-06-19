package roman

import (
	"errors"
)

// MarshalText encodes the receiver into UTF-8-encoded text and returns the result.
func (n *Roman) MarshalText() (text []byte, err error) {
	return nil, errors.New("not implemented")
}
