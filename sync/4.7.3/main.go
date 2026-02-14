package main

import (
	"fmt"
	"sync"
)

//Реализуйте тип Counter, который позволит нескольким горутинам безопасно работать с картой частот слов.
//
//Интерфейс:
//
//// увеличивает значение по ключу на 1
//Increment(str string)
//
//// возвращает значение по ключу
//Value(str string) int
//
//// проходит по всем записям, и для каждой вызывает функцию fn,
//// передавая в нее ключ и значение
//Range(fn func(key string, val int))

// начало решения

type Counter struct {
	data map[string]int
	lock sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[str]++
}

func (c *Counter) Value(str string) int {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.data[str] // todo: check data race
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for k, v := range c.data {
		fn(k, v)
	}
}

func NewCounter() *Counter {
	return &Counter{
		data: make(map[string]int),
		lock: sync.Mutex{},
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
