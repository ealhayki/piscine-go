package main

import (
	"fmt"
)

func StrLen(str string) int {
	return len(str)
}

func main() {
	str := "Ebrahim"

	fmt.Println(StrLen(str))
}
