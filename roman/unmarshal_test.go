package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnarshal(t *testing.T) {
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
			err := r.UnmarshalText([]byte(val))
			assert.NoError(t, err)
			assert.Equal(t, idx, r.Value)
		}
	})

	// TODO: Unmarshal negative
	// TODO: Unmarshal 10-20
	// TODO: Unmarshal 50-60
	// TODO: Unmarshal 1000-1010
}
