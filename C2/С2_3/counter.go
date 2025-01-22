package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := Counter{value: 0}
	wg.Add(1)
	go c.Increment()
	wg.Wait()
	fmt.Println(c.GetValue())
	wg.Add(1)
	c.Increment()
	wg.Wait()
	fmt.Println(c.GetValue())
	//fmt.Println("Test")
}

type Counter struct {
	value int // значение счетчика
	mu    sync.RWMutex
}
type Count interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

func (c *Counter) Increment() {
	//defer wg.Done()
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	return c.value
}
