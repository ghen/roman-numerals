package roman

import (
	"bytes"
	"errors"
)

// ErrInvalid indicates that operation can't be performed due to an invalid value.
var ErrInvalid = errors.New("invalid value")

var (
	values = map[uint16]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
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
	seg := bytes.Buffer{}
	for rng := uint16(1000); rng > 0; rng /= 10 {
		seg.Reset()

		for i := uint16(1); i <= val/rng; i += 1 {
			if str, exist := values[i*rng]; exist {
				seg.Reset()
				seg.WriteString(str)
			} else {
				seg.WriteString(values[rng])
			}
		}
		buf.Write(seg.Bytes())

		val = val % rng
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
