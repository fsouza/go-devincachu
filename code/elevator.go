package main

import (
	"fmt"
	"math/rand"
	"time"
)

func elevator(people chan string) {
	for person := range people {
		fmt.Printf("Carrying %s...\n", person)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

func main() {
	people := []string{
		"John", "Mary", "Gopher", "Glenda",
		"Jeff", "Paul", "Tux", "Pony",
	}
	ch := make(chan string)
	go elevator(ch)
	for _, p := range people {
		ch <- p
	}
	time.Sleep(1 * time.Second)
}
