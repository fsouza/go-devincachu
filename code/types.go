package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Birth time.Time
}

func (p *Person) Age() int {
	difference := time.Now().Sub(p.Birth)
	return int(difference / (365 * 24 * time.Hour))
}

type MyString string

type MyInt int

func main() {
	p := Person{
		Name:  "Francisco",
		Birth: time.Date(1989, 2, 16, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(p.Age())
}
