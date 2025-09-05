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
	if array[0]+array[1] > array[2] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
