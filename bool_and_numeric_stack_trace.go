package main

import (
	"fmt"
)

func main() {
	iPanic(true, 6, 'c', 6.7, complex(2, 1))
}

func iPanic(b bool, by byte, r rune, f float32, c complex64) {
	fmt.Println(by + 1)
	panic("I'm outta here")
}
