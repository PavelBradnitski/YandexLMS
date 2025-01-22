package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// //wg.Add(1)
	// go Send(ch1, ch2)
	// //wg.Wait()

	// vals1 := []int{}

	// for i := 0; i < 3; i++ {
	// 	val := <-ch1
	// 	vals1 = append(vals1, val)
	// }

	// //slices.Sort(vals1)
	// fmt.Println(vals1)

	// vals2 := []int{}

	// for i := 0; i < 3; i++ {
	// 	val := <-ch2
	// 	vals2 = append(vals2, val)
	// }

	//slices.Sort(vals2)
	// fmt.Println(vals2)

	ch := Process([]int{1, 2, 3})
	vals3 := []int{}
	fmt.Println(len(ch))
	for i := 0; i < 3; i++ {
		val := <-ch
		vals3 = append(vals3, val)
	}

	fmt.Println(vals3)
}

func SendOld(ch chan int, num int) {
	ch <- num
}

func Receive(ch chan int) int {
	return <-ch
}
func Send(ch1, ch2 chan int) {
	//defer wg.Done()
	for i := 0; i < 3; i++ {
		go SendOld(ch1, i)
		go SendOld(ch2, i)
	}
}

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for i := range nums {
		ch <- i
	}
	return ch
}
