package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "", NewRoman(0).String())

	assert.Equal(t, "IX", NewRoman(9).String())
	assert.Equal(t, "XIV", NewRoman(14).String())
	assert.Equal(t, "MCMXCIX", NewRoman(1999).String())
}
