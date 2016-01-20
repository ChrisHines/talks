package main

import (
	"bufio"
	"encoding/hex"
	"os"
)

func main() {
	buf := []byte{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		b := s.Bytes()
		if n := hex.EncodedLen(len(b)); len(buf) < n {
			buf = make([]byte, n)
		}
		n := hex.Encode(buf, b)
		_, err := os.Stdout.Write(buf[:n])
		panicOnError(err)
		_, err = os.Stdout.WriteString("\n")
		panicOnError(err)
	}
	panicOnError(s.Err())
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
