package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.False(t, NewRoman(0).IsValid())
	assert.False(t, NewRoman(4000).IsValid())

	assert.True(t, NewRoman(1).IsValid())
	assert.True(t, NewRoman(500).IsValid())
	assert.True(t, NewRoman(3999).IsValid())
}

func TestString(t *testing.T) {
	assert.Equal(t, "I", NewRoman(1).String())

	assert.Equal(t, "IX", NewRoman(9).String())
	assert.Equal(t, "XIV", NewRoman(14).String())
	assert.Equal(t, "MCMXCIX", NewRoman(1999).String())
}
