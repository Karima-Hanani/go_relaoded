package project

import (
	"strconv"
	"strings"
)

func Flag(line []string) []string {
	for i := 0; i < len(line); i++ {
		op := line[i]
		if op == "(cap)" || op == "(low)" || op == "(up)" || op == "(hex)" || op == "(bin)"{
			if i != 0 {
				switch op {
				case "(cap)":
					line[i-1] = Capitalize(line[i-1])
				case "(low)":
					line[i-1] = strings.ToLower(line[i-1])
				case "(up)":
					line[i-1] = strings.ToUpper(line[i-1])
				case "(hex)":
					n, err := strconv.ParseInt(line[i-1], 16, 64)
					if err != nil {
						line = append(line[:i], line[i+1:]...)
						i--
						continue
					} else {
						line[i-1] = strconv.Itoa(int(n))
					}
				case "(bin)":
					n, err := strconv.ParseInt(line[i-1], 2, 64)
					if err != nil {
						line = append(line[:i], line[i+1:]...)
						i--
						continue
					} else {
						line[i-1] = strconv.Itoa(int(n))
					}
				}
			}
			line = append(line[:i], line[i+1:]...)
			i--
			continue
		}
	}
	return line
}
