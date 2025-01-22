package main

import (
	"fmt"
	"slices"
)

func main() {
	sliceInt1 := []int{4, 1, 5, 0}
	sliceInt2 := []int{-1, 4, 5, 10}
	fmt.Println(SortAndMerge(sliceInt1, sliceInt2))
}

func SortAndMerge(left, right []int) []int {
	slices.SortFunc(left, SortInt)
	slices.SortFunc(right, SortInt)
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}

func SortInt(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
