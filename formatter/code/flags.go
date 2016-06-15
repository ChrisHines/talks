package main

import "fmt"

type twoStr struct {
	a, b string
}

func main() {
	fmt.Printf("%+d\n", 9)

	fmt.Printf("%+q\n", "Hello, 世界")

	hw := twoStr{"Hello", "world!"}
	fmt.Printf("%+v\n", hw)
}
