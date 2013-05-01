package main

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Birth time.Time
}

func NewPerson(name string, birth time.Time) (*Person, error) {
	if time.Now().Sub(birth) < 0 {
		return nil, errors.New("LOLWUT, did you born in the future?!")
	}
	person := Person{Name: name, Birth: birth}
	return &person, nil
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
