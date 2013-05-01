package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
	fmt.Println(sum(10, 5))
	fmt.Println(swap(10, 5))
	fmt.Println(split(20))
}
