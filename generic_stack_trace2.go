package main

func main() {
	iPanic(5)
}

func iPanic(i int) {
	if i > 0 {
		iPanic(i - 1)
	}
	panic("I'm outta here")
}
