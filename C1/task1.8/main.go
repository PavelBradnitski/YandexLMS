package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if match, _ := regexp.Match(`^[a-zA-Z]+$`, []byte(name)); !match {
			fmt.Fprintf(w, "Hello, dirty hacker!")
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Hello, stranger!")
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

// func StartServer(t time.Duration) {
// 	_ = t
// 	mux := http.NewServeMux()
// 	mux.Handle("/", SetDefaultName(Sanitize(HelloHandler)))
// 	http.ListenAndServe(":8080", mux)
// }

// func main() {
// 	StartServer(time.Millisecond)
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

var requestCount int
var mu sync.Mutex

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
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

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/metrics" {
			next.ServeHTTP(w, r)
		} else {
			FibonacciHandler(w, r)
		}
	}
}
func StartServer(t time.Duration) {
	_ = t
	http.HandleFunc("/", Metrics(MetricsHandler))
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer(time.Millisecond)
}
