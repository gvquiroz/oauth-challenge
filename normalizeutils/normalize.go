package normalizeutils

import (
	"strings"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"github.com/djimenez/iconv-go"
	"unicode/utf8"
)

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
	input := s;
	if(!utf8.ValidString(s)){
		input, _ = iconv.ConvertString(input, "latin1", "utf-8")
	}

	l := strings.ToLower(input)
	o := OmitAccent(l)
	q := RemoveUnwantedSymbols(o, "?")
	result := RemoveUnwantedSymbols(q, "!")

	return result
}