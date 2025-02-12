package main

import (
	"context"
	"io"
	"os"
	"time"
)

func main() {
	result := make(chan []byte)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	readJSON(ctx, "C:/golang/YandexLMS/C2/С2_7/C2_7_2/json1.json", result)

}
func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)
	file, _ := os.OpenFile(path, os.O_RDONLY, 0666)
	defer file.Close()
	data := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := file.Read(data)
			result <- data[:n]
			if err == io.EOF {
				break
			}

		}
	}
}
