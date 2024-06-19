// roman package defines types and interfaces to manipulate roman values.
package roman

// MinValue represents minimal supported roman number value.
const MinValue = uint16(1)

// MaxValue represents maximum supported roman number value.
const MaxValue = uint16(3999)

// Roman stores a roman number value.
type Roman struct {
	// Stores the value for the roman number.
	Value uint16
}

// NewRoman returns a new [Roman] instance for the [value].
func NewRoman(value uint16) Roman {
	return Roman{
		Value: value,
	}
}

// IsValid indicates if the [Roman.Value] can be represented in a roman notation.
func (r Roman) IsValid() bool {
	return r.Value >= MinValue && r.Value <= MaxValue
}

// String method is used to print values passed as an operand to any format
// that accepts a string or to an unformatted printer such as [fmt.Print].
func (r Roman) String() string {
	str, err := r.MarshalText()
	if err != nil {
		return err.Error()
	}
	return string(str)
}
