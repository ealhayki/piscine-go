package main

import (
	"fmt"
	"sort"
)

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	sort.Ints(a)

	return a[len(a)-1]
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}

	max := Max(a)

	fmt.Println(max)
}
