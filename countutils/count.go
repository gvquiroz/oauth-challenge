package countutils

import (
	"github.com/gvquiroz/oauth-challenge/normalizeutils"
)

func CountDuplicatedChars(s string) map[string]int {
	resultMap := make(map[string]int)

	// lower case - remove acents, question marks and exclamation marks
	normalizedString := normalizeutils.Normalize(s)

    for _, r := range normalizedString {
		char := string(r)

		if (resultMap[char] == 0) {
			resultMap[char] = 1
		} else {
			v := resultMap[char]
			resultMap[char] = v + 1
		}
    }

	return resultMap
}