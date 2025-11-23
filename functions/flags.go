package project

import (
	"strconv"
	"strings"
)

func Flag(line []string) []string {
	for i := 0; i < len(line); i++ {
		op := line[i]
		if op == "(cap)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
			}
			if i != 0 && i-1 >= 0 {
				line[i-1] = Capitalize(line[i-1])
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}

		if op == "(up)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}

			}
			if i != 0 && i-1 >= 0 {
				line[i-1] = strings.ToUpper(line[i-1])
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}

		if op == "(low)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}

			}
			if i != 0 && i-1 >= 0 {
				line[i-1] = strings.ToLower(line[i-1])
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}

		if op == "(bin)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}

			}
			if i != 0 && i-1 >= 0 {
				n, err := strconv.ParseInt(line[i-1], 2, 64)
				if err != nil {
					line = append(line[:i], line[i+1:]...)
					i--
					continue
				} else {
					line[i-1] = strconv.Itoa(int(n))
				}
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}

		if op == "(hex)" {
			for j := i - 1; j >= 0; j-- {
				if line[j] != "" {
					break
				}
			}
			if i != 0 {
				n, err := strconv.ParseInt(line[i-1], 16, 64)
				if err != nil {
					line = append(line[:i], line[i+1:]...)
					i--
					continue
				} else {
					line[i-1] = strconv.Itoa(int(n))
				}
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}
	}
	return line
}
