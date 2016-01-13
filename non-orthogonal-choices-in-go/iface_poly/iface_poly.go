package main

import (
	"bytes"
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
	data := []io.Reader{
		strings.NewReader("Hello, world!\n"),
		bytes.NewReader([]byte{71, 111, 111, 100, 98, 121, 101, 46, 10}),
		readerFunc(newReader("Once more unto the breach!\n")),
	}

	for _, r := range data {
		ioCopy(os.Stdout, r)
	}
}

func ioCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf) // HL
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr]) // HL
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
