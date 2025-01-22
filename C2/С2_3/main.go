package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func main() {
	runtime.GOMAXPROCS(10)

	concurrentMap := NewSafeMap()

	var wg sync.WaitGroup

	wg.Add(2000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			concurrentMap.Set(fmt.Sprintf("%d", rand.Intn(100)), rand.Intn(100))
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			val := concurrentMap.Get(fmt.Sprintf("%d", rand.Intn(100)))
			if val != nil {
				fmt.Println(val)
			} else {
				fmt.Println("Not found")
			}
		}()
	}

	wg.Wait()

}
func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}
func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.m[key]
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}
