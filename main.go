package main

import (
	"fmt"
	"os"
	"strings"

	project "project/functions"
)

func main() {
	// check the args
	if !project.CheckNumberOfArgs() {
		fmt.Println("Error: Number of arguments invalid !! ")
		return
	}
	if !project.CheckType(os.Args[1]) || !project.CheckType(os.Args[2]) {
		fmt.Println("Error: The arguments must be a .txt files.")
		return
	}

	input := project.Convert(strings.Split(project.ReadFile(os.Args[1]), "\n"))
	fmt.Println(input)

}
