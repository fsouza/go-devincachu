package main

import (
	"fmt"
)

func sample1() {
	var name string
	name = "Francisco"
	fmt.Println(name)
}

func sample2() {
	var name string = "Francisco"
	fmt.Println(name)
}

func sample3() {
	var name = "Francisco"
	fmt.Println(name)
}

func sample4() {
	name := "Francisco"
	fmt.Println(name)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Looping forever")
		break // not really
	}

	var f func(fn func(int) int) func(int) int
}
