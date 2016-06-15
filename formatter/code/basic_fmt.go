package main

import "fmt"

type twoStr struct {
	a, b string
}

func main() {
	fmt.Println("Hello, world!")

	hw := twoStr{"Hello", "world!"}
	fmt.Println(hw)
}
