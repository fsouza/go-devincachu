package main

import (
	"fmt"
)

func Reverse(values []int) []int {
	result := make([]int, len(values))
	length := len(values)
	for i := range values {
		result[i] = values[length-1-i]
	}
	return result
}

func main() {
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Reverse(numbers[:]))
}
