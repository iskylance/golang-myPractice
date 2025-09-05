package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	array := []int{a, b, c}
	sort.Ints(array)
	fmt.Print(array[1])
}
