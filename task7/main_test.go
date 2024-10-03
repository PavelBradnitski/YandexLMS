package main

import (
	"testing"
)

func TestPrintHello(t *testing.T) {
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}
}

func TestPrintHelloIvan(t *testing.T) {
	got := PrintHello("Ivan")
	expected := "Hello, Ivan!"

	if got != expected {
		t.Fatalf(`PrintHello("Ivan") = %q, want %q`, got, expected)
	}
}

func TestSum(t *testing.T) {
	got := Sum(2, 2)
	expected := 4
	if got != expected {
		t.Fatalf(`Sum(2,2) = %d, want %d`, got, expected)
	}
}

func TestLength(t *testing.T) {
	arrgot := []int{-1, 0, 1, 11, 111}
	arrwant := []string{"negative", "zero", "short", "long", "very long"}
	var got string
	for i, v := range arrgot {
		got = Length(v)
		if got != arrwant[i] {
			t.Fatalf(`Length(2) = %q, want %q`, got, arrwant[i])
		}
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(2, 2)
	expected := 4
	if got != expected {
		t.Fatalf(`Multiply(2,2) = %d, want %d`, got, expected)
	}
}

func TestDeleteVowels(t *testing.T) {
	s := "abcdeiou"
	got := DeleteVowels(s)
	expected := "bcd"
	if got != expected {
		t.Fatalf(`DeleteVowels(%s) = %q, want %q`, s, got, expected)
	}
}

func TestGetUTFLength(t *testing.T) {
	valid := []byte("Hello")
	got, _ := GetUTFLength(valid)
	expected := 5
	if got != expected {
		t.Fatalf(`GetUTFLength(%s) = %q, want %q`, valid, got, expected)
	}
	invalid := []byte{0xff, 0xfe, 0xfd}
	ab, err := GetUTFLength(invalid)
	_ = ab
	errExp := ErrInvalidUTF8
	if err != errExp {
		t.Fatalf(`GetUTFLength(%s) = %q, want %q`, invalid, err, errExp)
	}
}
