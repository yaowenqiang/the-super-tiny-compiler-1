package main

import "fmt"

func main() {
	iPanic("hello", make([]string, 3, 5))
}

func iPanic(s string, v []string) {
	fmt.Println(s + "World")
	panic("I'm outta here.")
}
