package main

import (
	"fmt"

	"github.com/go-stack/stack"
	"github.com/pkg/errors"
)

func main() {
	err := sub()
	fmt.Printf("%+v\n", err)
}

func sub() error {
	fmt.Printf("%+n\n\n", stack.Caller(0)) // go vet: "unrecognized printf verb 'n'"

	err := badCall()
	if err != nil {
		return err
	}
	return nil
}

func badCall() error {
	return errors.New("it's a trap")
}
