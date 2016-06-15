package main

import (
	"fmt"
	"io"
	"strings"
)

type twoStr struct {
	a, b string
}

func (s twoStr) Format(f fmt.State, verb rune) {
	io.WriteString(f, s.a+" ")

	w, ok := f.Width()
	if !ok {
		w = 1
	}
	if f.Flag('+') {
		w *= 2
	}
	io.WriteString(f, strings.Repeat(string(verb), w))

	io.WriteString(f, " "+s.b)
}

func main() {
	hw := twoStr{"Hello", "world!"}

	fmt.Printf("%5z\n", hw)  // go vet: "unrecognized printf verb 'z'"
	fmt.Printf("%+4&\n", hw) // go vet: "unrecognized printf verb '&'"
	fmt.Println(hw)
}
