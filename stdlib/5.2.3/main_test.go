package main

import (
	"testing"
)

func Test_prettify(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}
