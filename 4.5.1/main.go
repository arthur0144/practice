//Реализуйте тип Queue — очередь фиксированного размера на n элементов. Поддерживает две операции:
//
//Get(block bool) (int, error)
//Put(val int, block bool) error
//
//
//Put() помещает значение val в очередь:
//
//Если в очереди есть место — помещает val в очередь и возвращает nil.
//Если очередь заполнена и block = true — блокируется, пока не освободится место. Затем помещает val в очередь и возвращает nil.
//Если очередь заполнена и block = false — возвращает ошибку ErrFull.
//Get() выбирает значение из очереди:
//
//Если в очереди есть значения — выбирает очередное значение, возвращает его и nil.
//Если очередь пуста и block = true — блокируется, пока не появится значение. Затем выбирает и возвращает его и nil.
//Если очередь пуста и block = false — возвращает нулевое значение и ошибку ErrEmpty.
//Очередь работает по принципу FIFO (first in — first out): какое значение добавили раньше, такое раньше и вернется.
//Методы очереди могут вызываться из разных горутин.

package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// Queue - FIFO-очередь на n элементов
type Queue chan int

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if block {
		return <-q, nil
	}
	select {
	case v := <-q:
		return v, nil
	default:
		return 0, ErrEmpty
	}
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if block {
		q <- val
		return nil
	}
	select {
	case q <- val:
		return nil
	default:
		return ErrFull
	}
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return make(chan int, n)
}

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
