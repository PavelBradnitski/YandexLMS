package main

import (
	"fmt"
	"sync"
)

//var wg sync.WaitGroup

func main() {
	c := ConcurrentQueue{queue: make([]interface{}, 0)}
	// wg.Add(1)
	c.Enqueue("1")
	// wg.Wait()
	fmt.Println(c.Dequeue())
	// wg.Add(1)
	c.Enqueue("2")
	// wg.Wait()
	fmt.Println(c.Dequeue())
}

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
	queue []interface{} // здесь хранить элементы очереди
	mutex sync.Mutex
}

func (c *ConcurrentQueue) Enqueue(elem interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = append(c.queue, elem)
}
func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	out := c.queue[0]
	c.queue = c.queue[1:]
	return out
}
