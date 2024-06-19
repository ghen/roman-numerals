package roman

import (
	"errors"
)

// UnmarshalText decodes the UTF-8-encoded text and updates the receiver.
func (n *Roman) UnmarshalText(text []byte) error {
	return errors.New("not implemented")
}
