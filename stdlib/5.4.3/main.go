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
	for sc.Scan() {
		word := sc.Text()
		var f rune
		for _, f = range word {
			break
		}
		f = unicode.ToTitle(f)
		words = append(words, string(f)+word[len(string(f)):])
	}
	fmt.Println(strings.Join(words, " "))
}
