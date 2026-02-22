package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

var baseTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	nsec := t.UnixNano()
	s := strconv.FormatInt(nsec, 10)
	n := len(s)
	if n < 9 {
		i := "0"
		f := strings.TrimRight(s, "0")
		if f == "" {
			f = "0"
		}
		return i + "." + f
	}
	i := s[0 : n-9]
	f := strings.TrimRight(s[n-9:], "0")
	if f == "" {
		f = "0"
	}
	return i + "." + f
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	g := strings.Split(d, ".")
	if len(g) < 2 || g[0] == "" || g[1] == "" {
		return time.Time{}, fmt.Errorf("")
	}
	//fmt.Println(d)
	if len(g[1]) < 9 {
		d += strings.Repeat("0", 9-len(g[1]))
	}
	//fmt.Println("before cut", d)
	d = strings.ReplaceAll(d, ".", "")
	//fmt.Println("after cut", d)
	nsec, _ := strconv.ParseInt(d, 10, 64)
	return time.Unix(0, nsec).In(time.UTC), nil
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"

	input1 := "1970-01-01 01:00:00.000123456"
	t1, _ := time.Parse(layout, input1)

	input2 := "2022-05-24 14:45:22.951"
	t2, _ := time.Parse("2006-01-02 15:04:05.000", input2)

	input3 := "1970-01-01 01:00:00.000000001"
	t3, _ := time.Parse(layout, input3)

	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",

		t1: "3600.000123456",
		t2: "1653403522.951",
		t3: "3600.000000001",

		/*
			main.go:56: 1970-01-01 01:00:00.123456 +0000 UTC: got 3600.123456000, want 3600.123456
			main.go:56: 2022-05-24 14:45:22.951205 +0000 UTC: got 1653403522.951205000, want 1653403522.951205
			main.go:56: 1970-01-01 01:00:00.000000123 +0000 UTC: got 3600.123, want 3600.000000123*/
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Errorf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Errorf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Errorf("%v: got %v, want %v", src, got, want)
		}
	}
}
