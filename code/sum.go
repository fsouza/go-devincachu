package main

import (
	"fmt"
)

func sum(values []int, result chan int) {
	var sum int
	for _, v := range values {
		sum += v
	}
	result <- sum // HL
}

func main() {
	ch := make(chan int) // HL
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go sum(values, ch)
	fmt.Println(<-ch) // HL
}
