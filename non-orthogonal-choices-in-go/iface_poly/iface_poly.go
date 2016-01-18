package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readers := []io.Reader{
		strings.NewReader("Hello, world!\nHello, again!"),
		mustOpen("iface_poly/shakespeare.txt"),
	}

	for _, r := range readers {
		s := bufio.NewScanner(r) // scans lines of text
		for s.Scan() {
			fmt.Print("--", s.Text())
		}
		fmt.Println()
		if r, ok := r.(io.Closer); ok {
			r.Close() // close the reader if needed
		}
		if err := s.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// mustOpen returns an opened file or panics
func mustOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}
