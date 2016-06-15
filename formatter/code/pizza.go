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
	io.WriteString(f, strings.Repeat(string(verb), w))

	io.WriteString(f, " "+s.b)
}

func main() {
	hw := twoStr{"I'll take", "please."}

	fmt.Printf("%2🍕\n", hw) // go vet: "unrecognized printf verb '🍕'"
}
