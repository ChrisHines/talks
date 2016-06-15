package main

import (
	"errors"
	"fmt"
	"reflect"
)

type twoStr struct {
	a, b string
}

func (s twoStr) String() string {
	return s.a + " ... " + s.b
}

func main() {
	i := 5
	ri := reflect.ValueOf(i)
	fmt.Printf("%+d\n", ri) // go vet: "arg ri for printf verb %d of wrong type: reflect.Value"

	err := errors.New("oops you made a misteak")
	fmt.Printf("main failed: %v\n", err)

	hw := twoStr{"Hello", "world!"}
	fmt.Println(hw)
}
