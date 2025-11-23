package project

import (
	"fmt"
	"strconv"
)

func FlagWithPara(line []string) []string {
	tempstr := ""
	for i := 0; i < len(line); i++ {
		op := line[i]
		if op == "(up," {
			
			temp := line[i+1]
			for _, ch := range temp {
				if ch != ')' {
					tempstr += string(ch)
				}
			}
			para, _ := strconv.Atoi(tempstr)
			fmt.Println(para)
		}
		
	}
	return line
}
