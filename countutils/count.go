package countutils

import (
	"strings"
	"unicode"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/runes"
)

func CountChars(s string) map[string]int {
	m := make(map[string]int)

	// string process pipeline 
	// transform to lower case - remove acents, question marks and exclamation marks
	stringLowerCase := strings.ToLower(s)
	stringWithNoAcents := OmitAccent(stringLowerCase)
	stringWithoutQuestionMark := RemoveUnwantedSymbols(stringWithNoAcents, "?")
	stringToCount := RemoveUnwantedSymbols(stringWithoutQuestionMark, "!")

    for _, r := range stringToCount {
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