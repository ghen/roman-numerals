package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	var numerals = []string{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
	}

	t.Run("Marshal 0 (zero)", func(tt *testing.T) {
		r := Roman{Value: 0}
		b, err := r.MarshalText()
		require.NoError(t, err)
		assert.Equal(t, []byte(""), b)
	})

	t.Run("Marshal positive (0-9)", func(tt *testing.T) {
		for idx, val := range numerals {
			r := Roman{Value: idx}
			b, err := r.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, []byte(val), b)
		}
	})

	// TODO: Marshal negative
	// TODO: Marshal 10-20
	// TODO: Marshal 50-60
	// TODO: Marshal 1000-1010
}
