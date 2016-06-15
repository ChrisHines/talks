package main

import "fmt"

type twoStr struct {
	a, b string
}

func main() {
	fmt.Printf("%v\n", "Hello, world!")

	hw := twoStr{"Hello", "world!"}
	fmt.Printf("%v\n", hw)
}
