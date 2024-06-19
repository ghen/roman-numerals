// roman package defines types and interfaces to manipulate roman values.
package roman

// Roman stores a roman number value.
type Roman struct {
	// Stores the value for the roman number.
	Value int
}

// NewRoman returns a new [Roman] instance for the [value].
func NewRoman(value int) Roman {
	return Roman{
		Value: value,
	}
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
