// Аналог time.AfterFunc()
// Напишите функцию delay() — аналог time.AfterFunc():
//
// Принимает на входе интервал времени dur и функцию fn.
// Возвращает функцию отмены cancel.
// После запуска ждет в течение dur времени, после чего выполняет fn.
// Если клиент вызвал cancel() до истечения dur — не выполняет fn.
// Если клиент вызвал cancel() после истечения dur — ничего не делает.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cc := make(chan struct{})
	dc := make(chan struct{})

	go func() {
		select {
		case <-cc:
			cc = nil
			return
		default:
			time.Sleep(dur)
			dc <- struct{}{}
		}

		select {
		case <-dc:
			fn()
		case <-cc:
			cc = nil
			return
		}
	}()

	return func() {
		if cc == nil {
			return
		}
		cc <- struct{}{}
	}
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
