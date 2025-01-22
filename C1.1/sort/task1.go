package main

import (
	"fmt"
	"slices"
)

func main() {
	//slice123 := []uint{3, 2, 1}
	SliceString := []string{"Аксинья", "Варвара", "Арина", "Есения"}
	SortNames(SliceString)
	fmt.Println(SliceString)
}

func SortNums(nums []uint) {
	slices.SortFunc(nums, func(a, b uint) int {
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	})
}
func SortNames(names []string) {
	slices.SortFunc(names, func(a, b string) int {
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	})
}
