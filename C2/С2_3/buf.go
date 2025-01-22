package main

import (
	"fmt"
	"sync"
)

var Buf = []int{}
var mu sync.Mutex

func main() {

	// Write(1)
	// fmt.Println(Consume())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			Write(val)
		}(i)
	}

	wg.Wait()
	fmt.Println(Buf)
}

func Write(num int) {
	mu.Lock()
	defer mu.Unlock()
	Buf = append(Buf, num)
}
func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	out := Buf[0]
	Buf = Buf[1:]
	return out
}
