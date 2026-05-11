package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var words []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(nil, 3)
	for sc.Scan() {
		word := sc.Text()
		var f rune
		for _, f = range word {
			break
		}
		f = unicode.ToTitle(f)
		words = append(words, string(f)+word[len(string(f)):])
	}
	fmt.Println(len(words), sc.Err())
	fmt.Println(strings.Join(words, " "))
}
