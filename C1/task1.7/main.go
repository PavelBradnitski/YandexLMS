package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var requestCount int
var mu sync.Mutex

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/metrics" {
		FibonacciHandler(w, r)
	} else {
		MetricsHandler(w, r)
	}
}
func FibonacciHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%d", Fibonacci(requestCount))
	mu.Lock()
	requestCount++
	mu.Unlock()
}
func MetricsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func StartServer(t time.Duration) {
	_ = t
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer(time.Millisecond)
}
