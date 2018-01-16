package countutils

import (
	"strings"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func CountDuplicatedChars(s string) map[string]int {
	m := make(map[string]int)

	stringLowerCase := strings.ToLower(s)

	// remove acents, question marks and exclamation marks
	normalizedString := Normalize(stringLowerCase)

    for _, r := range normalizedString {
		c := string(r)

		if (m[c] == 0) {
			m[c] = 1
		} else {
			v := m[c]
			m[c] = v + 1
		}
    }

	return m
}

func OmitAccent(s string) string {
    t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
    r, _, _ := transform.String(t, s)
    return r
}

func RemoveUnwantedSymbols(s string, symbol string) string {
	t := strings.Replace(s, symbol, "", -1)
	return t
}

func Normalize(s string) string {

	stringWithNoAcents := OmitAccent(s)
	stringWithoutQuestionMarks := RemoveUnwantedSymbols(stringWithNoAcents, "?")
	stringResult := RemoveUnwantedSymbols(stringWithoutQuestionMarks, "!")

	return stringResult
}