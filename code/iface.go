package main

import (
	"fmt"
	"net/http"
	"os"
)

type Reader interface {
	Read(content []byte) (int, error)
}

func Dump(r Reader) {
	var buf [512]byte
	n, _ := r.Read(buf[:])
	for n > 0 {
		fmt.Printf("%s", buf)
		n, _ = r.Read(buf[:])
	}
}

func main() {
	resp, err := http.Get("http://golang.org/")
	if err != nil {
		panic(err)
	}
	Dump(resp.Body)
	resp.Body.Close()
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	Dump(file)
	file.Close()
}
