package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {

	// switch {
	// case num >= 0 && num < 10:
	// 	fmt.Println("Число меньше 10")
	// case num >= 10 && num < 100:
	// 	fmt.Println("Число меньше 100")
	// case num >= 100 && num < 1000:
	// 	fmt.Println("Число меньше 1000")
	// case num >= 1000:
	// 	fmt.Println("Число больше или равно 1000")
	// case num < 0:
	// 	fmt.Println("Некорректный ввод")
	// }
	var num int

	fmt.Scan(&num)
	FibonacciOutput(num)
}

func AreAnagrams(str1, str2 string) bool {
	if utf8.RuneCountInString(str1) != utf8.RuneCountInString(str2) {
		return false
	}
	str1 = strings.ToLower(str1)
	runes1 := []rune(str1)
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})

	str2 = strings.ToLower(str2)
	runes2 := []rune(str2)
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes1) == string(runes2)
}

func CalculateSeriesSum(n int) float64 {
	if n == 1 {
		return 1
	}
	return float64(1)/float64(n) + CalculateSeriesSum(n-1)
}

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	k1 = -k1
	k2 = -k2
	if k2 == k1 {
		return math.NaN(), math.NaN()
	}
	y := (k2*b1 - k1*b2) / (k2 - k1)
	x := (b2 - b1) / (k2 - k1)
	return x, y
}

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	newNum := 0
	for {
		if n/10 != 0 {
			newNum += n % 10
			n /= 10
		} else {
			newNum += n
			return CalculateDigitalRoot(newNum)
		}
	}
}

func IsPalindrome(input string) bool {
	str := strings.Join(strings.Split(input, " "), "")
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func FibonacciOutput(n int) {
	for i := 0; ; i++ {
		if n <= Fibonacci(i) {
			for j := 0; j < 10; j++ {
				fmt.Println(Fibonacci(i + j))
			}
			break
		}
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
