package main

import (
	"fmt"
	"slices"
	"strings"
)

// начало решения

// prettify возвращает отформатированное строковое представление карты
func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	}

	if len(m) == 1 {
		var k string
		var v int
		for k, v = range m {
		}
		return fmt.Sprintf("{ %s: %d }", k, v)
	}

	ss := make([]string, 0, len(m))
	template := "    %s: %d,"
	for k, v := range m {
		ss = append(ss, fmt.Sprintf(template, k, v))
	}
	slices.Sort(ss)

	return fmt.Sprintf("{\n%s\n}", strings.Join(ss, "\n"))
}

// конец решения
