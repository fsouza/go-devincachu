package main

import "fmt"

func say(what, who string) {
	fmt.Printf("%s, %s!\n", what, who)
}

func main() {
	say("Hello world", "#devincachu")
}
