package main

import (
	"fmt"
	"strings"
)

// начало решения

func isValid(c byte) bool {
	return 97 <= c && c <= 122 || 48 <= c && c <= 57 || c == 45 || 65 <= c && c <= 90
}

func slugify(src string) string {
	var b strings.Builder
	b.Grow(len(src))

	i, needSeparator := 0, false
	for ; i < len(src); i++ {
		c := src[i]
		if isValid(c) {
			if needSeparator {
				b.WriteByte(45)
				needSeparator = false
			}

			if 65 <= c && c <= 90 {
				c += 32
			}
			b.WriteByte(c)
			continue
		}
		needSeparator = true
	}

	return b.String()
}

func main() {
	const phrase = "Using Subtests and Sub-benchmarks !!!! -    "
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019
}
