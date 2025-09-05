package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)

	var array []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		array = append(array, num)
	}
	count := 0
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			count++
		}
	}
	fmt.Print(count)

}
