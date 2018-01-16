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

func TestCountDuplicatedCharactersWithLowerCase(t *testing.T) {
	m := CountChars("heLlo")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}

func TestCountDuplicatedCharactersWithLowerCaseOnMultipleChars(t *testing.T) {
	m := CountChars("HEllO")
	assert.Equal(t, 1, m["h"])
	assert.Equal(t, 2, m["l"])
	assert.Equal(t, 1, m["o"])
	assert.Equal(t, 1, m["e"])
}

func TestOmitAcents(t *testing.T) {
	m := OmitAccent("papá")
	assert.Equal(t, "papa", m)
}

func TestOmitDieresis(t *testing.T) {
	m := OmitAccent("päpa")
	assert.Equal(t, "papa", m)
}

func TestCountDuplicatedCharactersWithAcents(t *testing.T) {
	m := CountChars("papá")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])
}

func TestCountDuplicatedCharactersWithDieresis(t *testing.T) {
	m := CountChars("päpa")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])
}

func TestRemoveUnwantedSymbols(t *testing.T) {
	exclamationMark := "!"
	m := RemoveUnwantedSymbols("papa!",exclamationMark)
	assert.Equal(t, "papa", m)
}

func TestRemoveMultipleUnwantedSymbols(t *testing.T) {
	exclamationMark := "!"
	m := RemoveUnwantedSymbols("papa!!",exclamationMark)
	assert.Equal(t, "papa", m)
}

func TestRemoveUnwantedQuestionMark(t *testing.T) {
	questionMark := "?"
	m := RemoveUnwantedSymbols("papa?",questionMark)
	assert.Equal(t, "papa", m)
}

func TestCountDuplicatedCharactersWithExclamationMarks(t *testing.T) {
	m := CountChars("päpa!")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])

	assert.Equal(t, 0, m["!"])
}

func TestCountDuplicatedCharactersWithQuestionMarks(t *testing.T) {
	m := CountChars("päpa?")
	assert.Equal(t, 2, m["p"])
	assert.Equal(t, 2, m["a"])

	assert.Equal(t, 0, m["?"])
}
