package main

import (
	"bytes"
	"io"
	"strings"
)

func main() {
}

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}
func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024)
	n, err := r.Read(data)
	if err != io.EOF && err != nil {
		return "", err
	}
	return string(data[:n]), nil
}

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error) {
	u.UpperString = strings.ToUpper(string(p))
	return len(u.UpperString), err
}

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, n)
	nNew, err := r.Read(data)
	if err != io.EOF && err != nil {
		return err
	}
	_, err = w.Write(data[:nNew])
	return err
}

func Contains(r io.Reader, seq []byte) (bool, error) {
	data := make([]byte, 1024)
	_, err := r.Read(data)
	if err != io.EOF && err != nil {
		return false, err
	}
	return bytes.Contains(data, seq), nil
}
