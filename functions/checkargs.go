package project

import (
	"os"
	"strings"
)

func CheckNumberOfArgs() bool {
	if len(os.Args) != 3 {
		return false
	} else {
		return true
	}
}

func CheckType(str string) bool {
	if strings.HasSuffix(str, ".txt") {
		return true
	} else {
		return false
	}
}
