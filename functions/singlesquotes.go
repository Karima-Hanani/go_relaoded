package project

import (
	"unicode"
)

func SingleQuotes(s string) string {
	runes := []rune(s)
	out := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == ' ' {
			if len(out) > 0 && out[len(out)-1] == '\'' {
				j := i + 1
				for j < len(runes) && runes[j] == ' ' {
					j++
				}
				if j < len(runes) && IsPunc(runes[j]) {
					continue
				}
			}
			out = append(out, ' ')
			continue
		}

		if r != '\'' {
			out = append(out, r)
			continue
		}

		if IsApostrophe(runes, i) {
			out = append(out, '\'')
			continue
		}

		j := i + 1
		for j < len(runes) {
			if runes[j] == '\'' && !IsApostrophe(runes, j) {
				break
			}
			j++
		}
		if j == len(runes) { 
			out = append(out, '\'')
			continue
		}

		start, end := i+1, j
		for start < end && runes[start] == ' ' {
			start++
		}
		for end > start && runes[end-1] == ' ' {
			end--
		}

		if i > 0 && IsWordChar(runes[i-1]) && !IsPunc(runes[i-1]) &&
			(len(out) == 0 || out[len(out)-1] != ' ') {
			out = append(out, ' ')
		}

		out = append(out, '\'')
		out = append(out, runes[start:end]...)
		out = append(out, '\'')

		i = j

		if i+1 < len(runes) && IsWordChar(runes[i+1]) && !IsPunc(runes[i+1]) {
			out = append(out, ' ')
		}
	}

	return string(out)
}

func IsApostrophe(runes []rune, i int) bool {
	return i > 0 && i+1 < len(runes) && IsWordChar(runes[i-1]) && IsWordChar(runes[i+1])
}

func IsWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func IsPunc(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ';' || r == ':'
}
