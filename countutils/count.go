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
	stringLowerCase := strings.ToLower(s)
	stringWithNoAcents := OmitAccent(stringLowerCase)
    for _, r := range stringWithNoAcents {
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