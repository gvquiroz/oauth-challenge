package countutils

import (
	"strings"
	"github.com/gvquiroz/oauth-challenge/normalizeutils"
)

func CountDuplicatedChars(s string) map[string]int {
	m := make(map[string]int)

	stringLowerCase := strings.ToLower(s)

	// remove acents, question marks and exclamation marks
	normalizedString := normalizeutils.Normalize(stringLowerCase)

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