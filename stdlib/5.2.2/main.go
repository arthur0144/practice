package main

import (
	"strconv"
	"strings"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	var total float64
	for _, step := range directions {
		words := strings.Split(step, " ")
		for _, word := range words {
			if strings.HasSuffix(word, "km") {
				distance, err := strconv.ParseFloat(word[:len(word)-2], 64)
				if err != nil {
					continue
				}
				total += distance * 1000
				break
			}
			if strings.HasSuffix(word, "m") {
				distance, err := strconv.ParseFloat(word[:len(word)-1], 64)
				if err != nil {
					continue
				}
				total += distance
				break
			}
		}
	}
	return int(total)
}

// конец решения

func main() {

}
