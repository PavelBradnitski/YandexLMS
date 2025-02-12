package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(TimeoutFibonacci(100, time.Microsecond))
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int, 1)
	go func() {
		num := Fibonacci(n)
		c <- num
	}()
	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
