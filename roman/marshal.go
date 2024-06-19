package roman

import (
	"bytes"
	"errors"
)

// ErrInvalid indicates that operation can't be performed due to an invalid value.
var ErrInvalid = errors.New("invalid value")

var (
	alphabet = map[uint16]string{
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
	values = map[string]uint16{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
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
			if str, exist := alphabet[i*rng]; exist {
				seg.Reset()
				seg.WriteString(str)
			} else {
				seg.WriteString(alphabet[rng])
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
	return errors.New("not implemented")
}
