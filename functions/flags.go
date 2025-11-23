package project

import (
	"strconv"
	"strings"
)

func Flag(line []string) []string {
	count := 1
	for i := 0; i < len(line); i++ {
		op := line[i]
		if op == "(cap)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
				count++
			}
			if i != 0 && i-count >= 0 {

				line[i-count] = strings.Title(line[i-count])
			}
			line = append(line[:i], line[i+1:]...)
			count = 1
			i--
			continue
		}

		if op == "(up)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
				count++
			}
			if i != 0 && i-count >= 0 {
				line[i-count] = strings.ToUpper(line[i-count])
			}
			line = append(line[:i], line[i+1:]...)
			count = 1
			i--
			continue
		}

		if op == "(low)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
				count++
			}
			if i != 0 && i-count >= 0 {
				line[i-count] = strings.ToLower(line[i-count])
			}
			line = append(line[:i], line[i+1:]...)
			count = 1
			i--
			continue
		}

		if op == "(bin)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
				count++
			}
			if i != 0 && i-count >= 0 {
				n, err := strconv.ParseInt(line[i-count], 2, 64)
				if err != nil {
					line = append(line[:i], line[i+1:]...)
					i--
					continue
				} else {
					line[i-count] = strconv.Itoa(int(n))
				}
			}
			line = append(line[:i], line[i+1:]...)
			count = 1
			i--
			continue
		}

		if op == "(hex)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
				count++
			}
			if i != 0 {
				n, err := strconv.ParseInt(line[i-count], 16, 64)
				if err != nil {
					line = append(line[:i], line[i+1:]...)
					i--
					continue
				} else {
					line[i-count] = strconv.Itoa(int(n))
				}
			}
			line = append(line[:i], line[i+1:]...)
			count++
			i--
			continue
		}
	}

	return line
}
