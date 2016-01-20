package main

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	const file32 = "dynamic_behavior/GoSpec-base32.txt"
	c := CountEncodedWords(readLines(file32), "go", base32.StdEncoding) // HL
	fmt.Println(c, file32)

	const file64 = "dynamic_behavior/GoSpec-base64.txt"
	c = CountEncodedWords(readLines(file64), "go", base64.StdEncoding) // HL
	fmt.Println(c, file64)

	const filehex = "dynamic_behavior/GoSpec-hex.txt"
	c = CountEncodedWords(readLines(filehex), "go", StringDecoderFunc(hex.DecodeString)) // HL
	fmt.Println(c, filehex)

	const fileraw = "/Go/doc/go_spec.html"
	c = CountEncodedWords(readLines(fileraw), "go", NopStringDecoder{}) // HL
	fmt.Println(c, fileraw)
}

type NopStringDecoder struct{}

func (NopStringDecoder) DecodeString(s string) ([]byte, error) {
	return []byte(s), nil
}

type StringDecoderFunc func(s string) ([]byte, error)

func (fn StringDecoderFunc) DecodeString(s string) ([]byte, error) {
	return fn(s)
}

type StringDecoder interface { // HL
	DecodeString(s string) ([]byte, error) // HL
} // HL

// CountEncodedWords counts the occurances of word in messages. Messages are
// decoded with decoder.
func CountEncodedWords(messages []string, word string, decoder StringDecoder) int { // HL
	count := 0
	for _, m := range messages {
		if dm, err := decoder.DecodeString(m); err == nil { // HL
			count += bytes.Count(dm, []byte(word))
		}
	}
	return count
}

func readLines(path string) []string {
	f, err := os.Open(path)
	panicOnError(err)
	defer f.Close()

	var lines []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	panicOnError(s.Err())
	return lines
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
