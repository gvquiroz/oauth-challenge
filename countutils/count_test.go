package countutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountDuplicatedCharacters(t *testing.T) {
	m := CountChars("hello")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}