package main

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N%2 == 0 {
		IsPowerOfTwoRecursive(N / 2)
	} else if N == 1 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
func main() {
	IsPowerOfTwoRecursive(15)
}
