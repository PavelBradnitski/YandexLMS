package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic Test 1",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Basic Test 2",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Basic Test 3",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		SortIntegers(tt.input)
		got := tt.input
		if slices.Compare(got, tt.expected) != 0 {
			t.Errorf("SortIntegers() = %v, expected %v", got, tt.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected bool
	}{
		{
			name:     "Basic Test 1",
			input:    []int{1, 2, 3},
			target:   1,
			expected: true,
		},
		{
			name:     "Basic Test 2",
			input:    []int{1, 2, 3},
			target:   2,
			expected: true,
		},
		{
			name:     "Basic Test 3",
			input:    []int{1, 2, 3},
			target:   4,
			expected: false,
		},
	}

	for _, tt := range tests {
		got := Contains(tt.input, tt.target)
		if got != tt.expected {
			t.Errorf("Contains() = %v, expected %v", got, tt.expected)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic Test 1",
			input:    "abc",
			expected: "cba",
		},
		{
			name:     "Basic Test 2",
			input:    "bbg",
			expected: "gbb",
		},
		{
			name:     "Basic Test 3",
			input:    "luck",
			expected: "kcul",
		},
	}

	for _, tt := range tests {
		got := ReverseString(tt.input)
		if got != tt.expected {
			t.Errorf("ReverseString() = %v, expected %v", got, tt.expected)
		}
	}
}

func TestAreAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{
			name:     "Basic Test 1",
			input1:   "listen",
			input2:   "silent",
			expected: true,
		},
		{
			name:     "Basic Test 1",
			input1:   "The Morse Code",
			input2:   "Here come dots",
			expected: true,
		},
		{
			name:     "Basic Test 1",
			input1:   "a b",
			input2:   "a c",
			expected: false,
		},
		{
			name:     "Basic Test 4",
			input1:   "a b",
			input2:   "a ca",
			expected: false,
		},
	}

	for _, tt := range tests {
		got := AreAnagrams(tt.input1, tt.input2)
		if got != tt.expected {
			t.Errorf("AreAnagrams() = %v, expected %v", got, tt.expected)
		}
	}
}
