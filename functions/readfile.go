package project

import (
	"fmt"
	"os"
)

func ReadFile(str string) string {
	file, err := os.Open(str)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	input := ""
	for err == nil {
		input += string(buf[:n])
		n, err = file.Read(buf)
	}
	return input
}
