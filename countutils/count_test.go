package countutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountDuplicatedCharacters(t *testing.T) {
	m := CountDuplicatedChars("hello")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}

func TestCountDuplicatedCharactersWithLowerCase(t *testing.T) {
	m := CountDuplicatedChars("heLlo")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}

func TestCountDuplicatedCharactersWithLowerCaseOnMultipleChars(t *testing.T) {
	m := CountDuplicatedChars("HEllO")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}

func TestCountDuplicatedCharactersWithAcents(t *testing.T) {
	m := CountDuplicatedChars("papá")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])
}

func TestCountDuplicatedCharactersWithDieresis(t *testing.T) {
	m := CountDuplicatedChars("päpa")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])
}

func TestCountDuplicatedCharactersWithExclamationMarks(t *testing.T) {
	m := CountDuplicatedChars("päpa!")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])

	assert.Equal(t, 0, m["!"])
}

func TestCountDuplicatedCharactersWithQuestionMarks(t *testing.T) {
	m := CountDuplicatedChars("päpa?")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])

	assert.Equal(t, 0, m["?"])
}