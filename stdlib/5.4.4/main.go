package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

// начало решения

type RReader struct {
	buf []byte
	i   int
	max int
}

func (r *RReader) Read(b []byte) (n int, err error) {
	if r.i >= r.max {
		return 0, io.EOF
	}
	n = copy(b, r.buf[r.i:])
	r.i += n
	return
}

func RandomReader(max int) io.Reader {
	buf := make([]byte, max)
	rand.Read(buf)
	return &RReader{
		buf: buf,
		i:   0,
		max: max,
	}
}

// конец решения

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
		fmt.Println()
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}
