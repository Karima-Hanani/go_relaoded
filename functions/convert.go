package project

import (
	"fmt"
	"strings"
)

func Convert(input []string) [][]string {
	matrice := [][]string{}
	for _, lines := range input {
		line := strings.Split(lines, " ")
		cleanLine:= []string{}
		for _,ch := range line {
			if ch == ""{
				continue
			} 
			cleanLine = append(cleanLine, ch)
		}
		cleanLine = Flag(cleanLine)
		cleanLine = FlagWithPara(cleanLine)
		matrice = append(matrice, cleanLine)

	}
	fmt.Printf("%#v\n", matrice)
	return matrice
}
