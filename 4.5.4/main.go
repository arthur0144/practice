package main

import (
	"fmt"
	"time"
)

// начало решения
//Напишите функцию schedule(), которая запускает регулярное выполнение задачи:
//
//Принимает на входе интервал времени dur и функцию fn.
//Возвращает функцию отмены cancel.
//После запуска каждые dur времени выполняет fn.
//Если клиент вызвал cancel() — перестает выполнять fn.

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	cc := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fn()
			case <-cc:
				return
			}
		}
	}()

	return func() {
		select {
		case <-cc:
		default:
			close(cc)
		}
	}
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
