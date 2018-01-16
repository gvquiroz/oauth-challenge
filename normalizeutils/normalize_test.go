package normalizeutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeString(t *testing.T) {
	m := Normalize("päpa!?")
	assert.Equal(t, "papa", m)
}

func TestOmitAcents(t *testing.T) {
	m := OmitAccent("papá")
	assert.Equal(t, "papa", m)
}

func TestOmitDieresis(t *testing.T) {
	m := OmitAccent("päpa")
	assert.Equal(t, "papa", m)
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