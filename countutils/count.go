package countutils

import "strings"

func CountChars(s string) map[string]int {
	m := make(map[string]int)

	str := strings.ToLower(s)
    for _, r := range str {
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