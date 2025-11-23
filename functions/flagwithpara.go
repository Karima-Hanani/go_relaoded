package project

import (
	"strconv"
	"strings"
)

func FlagWithPara(line []string) []string {
	tempstr := ""
	for i := 0; i < len(line); i++ {
		op := line[i]
		if op == "(up," || op == "(cap," || op == "(low," {
			temp := line[i+1]
			if strings.HasSuffix(temp, ")") {
				for _, ch := range temp {
					if ch != ')' {
						tempstr += string(ch)
					}
				}
				para, err := strconv.Atoi(tempstr)
				if err != nil {
					break
				}
				for k := para; k > 0; k-- {
					if i-k >= 0 {
						switch op {
						case "(up,":
							line[i-k] = strings.ToUpper(line[i-k])
						case "(low,":
							line[i-k] = strings.ToLower(line[i-k])
						case "(cap,":
							line[i-k] = Capitalize(line[i-k])
						}
					}
				}
			} else {
				break
			}
			line = append(line[:i], line[i+2:]...)
		}
	}
	return line
}
