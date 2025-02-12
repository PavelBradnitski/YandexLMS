package main

import (
	"fmt"
	"math/rand"
	"time"
)

// функция генерирует случайное число в интервале [0, 100)
func random() int {
	const max int = 100
	return rand.Intn(max)
}

func main() {
	const size int = 10
	results := []int{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func() {
			results = append(results, random())
		}()
	}
	time.Sleep(time.Second)

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}
}
