package main

import (
	"errors"
	"fmt"
	"sort"
	"unicode"
)

func DeleteLongKeys(m map[string]int) map[string]int {
	for k := range m {
		if len(k) < 6 {
			delete(m, k)
		}
	}
	return m
}
func CountingSort(contacts []string) map[string]int {
	out := make(map[string]int, len(contacts))
	for _, v := range contacts {
		out[v]++
	}
	return out
}
func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := make(map[string]string, len(m))
	for k, v := range m {
		newMap[v] = k
	}
	return newMap
}

func SumOfValuesInMap(m map[int]int) int {
	out := 0
	for _, v := range m {
		out += v
	}
	return out
}
func FindMaxKey(m map[int]int) int {
	out := 0
	for k, v := range m {
		if v > out {
			out = k
		}
	}
	return out
}
func isLatin(input string) bool {
	for _, v := range input {
		if !unicode.In(v, unicode.Latin) {
			return false
		}
	}
	return true
}
func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}

func StringLength(input string) int {
	return len(input)
}
func Mix(nums []int) []int {
	split := len(nums) / 2
	var out []int
	if split <= 1 {
		return nums
	}

	for i := 0; i < split; i++ {
		out = append(out, nums[i], nums[i+split])
	}
	return out
}
func Join(nums1, nums2 []int) []int {
	out := make([]int, len(nums1)+len(nums2))
	copy(out, nums1)
	copy(out[len(nums1):], nums2)
	return out
}
func SliceCopy(nums []int) []int {
	newNum := make([]int, len(nums))
	copy(newNum, nums)
	return newNum
}
func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
			i--
			continue

		}
	}
	return nums
}
func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil || n < 0 {
		return nil, errors.New("FAIL")
	}
	var out []int
	count := 0
	for _, v := range nums {
		if v < limit && count < n {
			out = append(out, v)
			count++
		}
	}
	return out, nil
}
func PrettyArrayOutput(array [9]string) {
	var str string
	for i, v := range array {
		if i < 7 {
			str += fmt.Sprintf("%d я уже сделал: %s\n", i+1, v)
		} else if i >= 7 && i < 9 {
			str += fmt.Sprintf("%d не успел сделать: %s\n", i+1, v)
		}
	}
	fmt.Print(str)
}
func SumOfArray(array [6]int) int {
	var sum int
	for _, v := range array {
		sum += v
	}
	return sum
}
func FindMinMaxInArray(array [10]int) (int, int) {
	min := array[0]
	max := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max, min
}
func FiveSteps(array [5]int) [5]int {
	var slice []int
	slice = append(slice, array[:]...)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j] // Compare in descending order
	})
	var out [5]int
	copy(out[:], slice)
	return out
}

func main() {
	fmt.Println(DeleteLongKeys(map[string]int{"abcde": 1, "abcdef": 2}))
}
