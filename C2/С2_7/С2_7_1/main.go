package main

import (
	"fmt"
)

func main() {
	stop := make(chan struct{})
	primeChan := make(chan int)

	// test for generating primes up to 10
	//expectedPrimesUpTo10 := []int{2, 3, 5, 7}
	go GeneratePrimeNumbers(stop, primeChan, 1000)
	count := 0
	receivedPrimes := make([]int, 0)
	for prime := range primeChan {
		receivedPrimes = append(receivedPrimes, prime)
		count++
		if count == 3 {
			close(stop)
		}
	}
	fmt.Println(receivedPrimes)
	// close stop channel to terminate generatePrimeNumbers goroutine

}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= N; p++ {
		if isPrime[p] {
			for i := p * p; i <= N; i += p {
				isPrime[i] = false
			}
		}
	}
K:
	for p := 2; p <= N; p++ {
		if isPrime[p] {
		J:
			for {
				select {
				case <-stop:
					break K
				default:
					prime_nums <- p
					break J
				}
			}
		}
	}
	close(prime_nums)
}
