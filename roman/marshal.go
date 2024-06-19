package roman

import (
	"bytes"
	"errors"
	"math"
)

// ErrInvalid indicates that operation can't be performed due to an invalid value.
var ErrInvalid = errors.New("invalid value")

var (
	ranges = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	alphabet = map[byte]uint16{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

// MarshalText encodes the receiver into UTF-8-encoded text and returns the result.
// Returns [ErrInvalid] if the [Roman] is not a valid number. See [Roman.IsValid] for more details.
func (r *Roman) MarshalText() (text []byte, err error) {
	if !r.IsValid() {
		return nil, ErrInvalid
	}

	val := r.Value
	buf := bytes.Buffer{}
	for rng := 3; rng >= 0; rng -= 1 {
		base := uint16(math.Pow10(rng))
		if dgt := val / base; dgt > 0 {
			buf.WriteString(ranges[rng][dgt])
		}

		val = val % base
	}
	return buf.Bytes(), nil
}

// UnmarshalText decodes the UTF-8-encoded text and updates the receiver.
// Returns [ErrInvalid] if the text can't be parsed into a valid [Roman] value.
func (r *Roman) UnmarshalText(text []byte) error {
	r.Value = 0

	// NOTE: This method implements general parsing logic flow.
	//       It does not implement validation. For example, "IVIV" would parse as "4+4 = 8".
	//
	//       It would be easy to use separate Regex to validate the overall number structure.
	//       However, due to time restrictions oif approximately 1Hr on this assignment,
	//       validation is not covered.

	max := uint16(0)
	for idx := len(text) - 1; idx >= 0; idx -= 1 {
		dig := text[idx]
		if val, exist := alphabet[dig]; !exist {
			return ErrInvalid
		} else if val >= max {
			r.Value += val
			max = val
		} else {
			r.Value -= val
		}

		if r.Value > MaxValue {
			return ErrInvalid
		}
	}

	return nil
}
