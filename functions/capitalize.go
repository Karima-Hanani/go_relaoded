package project

import (
	"unicode"
)

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
