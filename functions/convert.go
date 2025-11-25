package project

import (
	"fmt"
	"strings"
)

func Convert(input []string) string {
	matrice := [][]string{}
	tempstr := ""
	for _, lines := range input {
		line := strings.Split(lines, " ")
		cleanLine := []string{}
		for _, ch := range line {
			if ch == "" {
				continue
			}
			cleanLine = append(cleanLine, ch)
		}
		cleanLine = Flag(cleanLine)
		cleanLine = FlagWithPara(cleanLine)
		matrice = append(matrice, cleanLine)

	}
	for i := 0 ; i < len(matrice) ; i ++ {
		tempstr += Ponctuations(matrice[i])
	}
	fmt.Printf("%#v\n", matrice)
	return tempstr
}
