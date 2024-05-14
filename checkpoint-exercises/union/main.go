package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		result := ""
		arg1 := os.Args[1]
		arg2 := os.Args[2]

		for _, ch := range arg1 {
			if !strings.ContainsRune(result, ch) {
				result += string(ch)
			}
		}

		for _, ch := range arg2 {
			if !strings.ContainsRune(result, ch) {
				result += string(ch)
			}
		}

		for _, ch := range result {
			z01.PrintRune(ch)
		}

		z01.PrintRune('\n')
	}
}
