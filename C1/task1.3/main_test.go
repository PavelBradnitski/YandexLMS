package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type customReader struct {
}

func (cr *customReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func NewCustomReader() *customReader {
	return &customReader{}
}
func TestReadString(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid Empty input",
			input:    strings.NewReader(""),
			expected: "",
			wantErr:  false,
		},
		{
			name:     "Valid input",
			input:    strings.NewReader("Hello, World!"),
			expected: "Hello, World!",
			wantErr:  false,
		},
		{
			name:    "Invalid reader",
			input:   NewCustomReader(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := ReadString(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReadString() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if got != tt.expected {
			t.Errorf("ReadString() = %v, expected %v", got, tt.expected)
		}
	}
}
