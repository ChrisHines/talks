package main

import (
	"io"
	"os"
	"strings"
)

func newReader(s string) func(p []byte) (n int, err error) {
	return func(p []byte) (n int, err error) {
		n = copy(p, s)
		s = s[n:]
		if len(s) == 0 {
			err = io.EOF
		}
		return n, err
	}
}

type readerFunc func(p []byte) (n int, err error)

func (f readerFunc) Read(p []byte) (n int, err error) {
	return f(p)
}

func main() {
	data := []func(p []byte) (n int, err error){
		newReader("Hello, world!\n"),
		strings.NewReader("Goodbye.\n").Read,
		newReader("Once more unto the breach!\n"),
	}

	for _, r := range data {
		ioCopy(os.Stdout.Write, r)
	}
}

func ioCopy(dst, src func(p []byte) (n int, err error)) (written int64, err error) {
	buf := make([]byte, 5)
	for {
		nr, er := src(buf) // HL
		if nr > 0 {
			nw, ew := dst(buf[0:nr]) // HL
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
