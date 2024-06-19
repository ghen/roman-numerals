package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalText(t *testing.T) {

	t.Run("Marshal 0 (zero)", func(tt *testing.T) {
		r := Roman{Value: 0}
		_, err := r.MarshalText()
		assert.Error(t, err, ErrInvalid)
	})

	t.Run("Marshal 4000", func(tt *testing.T) {
		r := Roman{Value: 4000}
		_, err := r.MarshalText()
		assert.Error(t, err, ErrInvalid)
	})

	t.Run("Marshal values", func(tt *testing.T) {
		for key, val := range values {
			r := NewRoman(key)

			b, err := r.MarshalText()
			require.NoError(t, err)

			assert.Equal(t, val, string(b))
		}
	})

	t.Run("Marshal random", func(tt *testing.T) {
		tests := map[uint16]string{
			11:   "XI",
			15:   "XV",
			27:   "XXVII",
			1999: "MCMXCIX",
		}

		for key, val := range tests {
			r := NewRoman(key)

			b, err := r.MarshalText()
			require.NoError(t, err)

			assert.Equal(t, val, string(b))
		}
	})
}

func TestUnmarshalText(t *testing.T) {

	t.Run("Unmarshal empty string", func(tt *testing.T) {
		r := NewRoman(0)

		err := r.UnmarshalText([]byte(""))
		assert.NoError(t, err)

		assert.Equal(t, uint16(0), r.Value)
	})

	t.Run("Unmarshal out of range", func(tt *testing.T) {
		r := NewRoman(0)
		err := r.UnmarshalText([]byte("MMMM"))
		assert.Error(t, err, ErrInvalid)
	})

	t.Run("Unmarshal alphabet", func(tt *testing.T) {
		for chr, val := range alphabet {
			r := NewRoman(0)

			err := r.UnmarshalText([]byte{chr})
			require.NoError(t, err)

			assert.Equal(t, val, r.Value)
		}
	})

	t.Run("Unmarshal random", func(tt *testing.T) {
		tests := map[string]uint16{
			"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
			"XI": 11, "XIV": 14, "XIIII": 14, "MCMXCIX": 1999,
		}

		for str, val := range tests {
			r := NewRoman(0)

			err := r.UnmarshalText([]byte(str))
			require.NoError(t, err)

			assert.Equal(t, val, r.Value)
		}
	})

}
