// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	cc := make(chan struct{})

	dur := time.Second / time.Duration(limit)
	ticker := time.NewTicker(dur)

	handle = func() error {
		select {
		case <-ticker.C:
			go fn()
		case _, ok := <-cc:
			if !ok {
				return ErrCanceled
			}
		}
		return nil
	}

	cancel = func() {
		select {
		case <-cc:
			return
		default:
			ticker.Stop()
			close(cc)
		}
	}

	return
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(10, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
