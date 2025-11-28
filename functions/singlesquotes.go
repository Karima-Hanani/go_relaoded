package project

func IsLetter(c rune) int {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return 1
	}
	return 0
}

func IsQuotePosition(b []rune, i int) int {
	if i > 0 && i < len(b)-1 {
		if IsLetter(b[i-1]) == 1 && IsLetter(b[i+1]) == 1 {
			return 0
		}
	}
	return 1
}

func RemoveSpacesInsideQuotes(s string) string {
	b := []rune(s)

	var quotes []int
	for i := 0; i < len(b); i++ {
		if b[i] == '\'' && IsQuotePosition(b, i) == 1 {
			quotes = append(quotes, i)
		}
	}

	if len(quotes) < 2 {
		return s
	}

	offset := 0

	for i := 0; i+1 < len(quotes); i += 2 {
		start := quotes[i] - offset
		end := quotes[i+1] - offset

		left := start + 1
		right := end - 1

		leadingSpaces := 0
		for left <= right && b[left] == ' ' {
			left++
			leadingSpaces++
		}

		trailingSpaces := 0
		for right >= left && b[right] == ' ' {
			right--
			trailingSpaces++
		}

		totalSpacesToRemove := leadingSpaces + trailingSpaces

		if totalSpacesToRemove > 0 {
			j := start + 1
			for k := left; k <= right; k++ {
				b[j] = b[k]
				j++
			}

			b[j] = '\''
			j++

			for k := end + 1; k < len(b); k++ {
				b[j] = b[k]
				j++
			}

			for j < len(b) {
				b[j] = ' '
				j++
			}

			offset += totalSpacesToRemove
		}
	}

	return string(b)
}
