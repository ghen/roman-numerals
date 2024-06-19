package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlphabet(t *testing.T) {
	for key, val := range alphabet {
		assert.Equal(t, key, values[val])
	}
}

func TestMarshalText(t *testing.T) {
	var numerals = []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
	}

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

	t.Run("Marshal positive (0-9)", func(tt *testing.T) {
		for idx, val := range numerals {
			r := NewRoman(uint16(idx + 1))

			b, err := r.MarshalText()
			require.NoError(t, err)

			assert.Equal(t, val, string(b))
		}
	})

	// TODO: Marshal 10-20
	// TODO: Marshal 50-60
	// TODO: Marshal 1000-1010
}

func TestUnmarshalText(t *testing.T) {
	var numerals = []string{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
	}

	t.Run("Unmarshal empty string", func(tt *testing.T) {
		r := NewRoman(0)
		err := r.UnmarshalText([]byte(""))
		assert.Error(t, err, ErrInvalid)
	})

	t.Run("Unmarshal out of range", func(tt *testing.T) {
		r := NewRoman(0)
		err := r.UnmarshalText([]byte("MMMM"))
		assert.Error(t, err, ErrInvalid)
	})

	t.Run("Marshal positive (1-9)", func(tt *testing.T) {
		for idx, val := range numerals {
			r := NewRoman(uint16(idx + 1))

			err := r.UnmarshalText([]byte(val))
			require.NoError(t, err)

			assert.Equal(t, idx, r.Value)
		}
	})

	// TODO: Unmarshal 10-20
	// TODO: Unmarshal 50-60
	// TODO: Unmarshal 1000-1010
}
