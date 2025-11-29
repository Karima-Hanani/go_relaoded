package project

import (
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
	for i := 0; i < len(matrice); i++ {
		if i != len(matrice)-1 {
			tempstr += Ponctuations(matrice[i]) + "\n"

		} else {
			tempstr += Ponctuations(matrice[i])
		}
	}
	return tempstr
}
