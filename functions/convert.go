package project

import (
	"fmt"
	"strings"
)

func Convert(input []string) [][]string {
	matrice := [][]string{}
	for _, lines := range input {
		line := Flag(strings.Split(lines, " "))
		matrice = append(matrice, line)

	}
	fmt.Printf("%#v\n", matrice)
	return matrice
}
