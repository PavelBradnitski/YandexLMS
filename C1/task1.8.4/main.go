package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if match, _ := regexp.Match(`^[a-zA-Z]+$`, []byte(name)); !match {
			panic("invalid name")
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			out := "{\"greetings\":\"hello\",\"name\":\"stranger\"}"
			w.Header().Set("myKey", out)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(w, "{\"status\":\"error\",\"result\":{}}")
			} else {
				value := w.Header().Get("myKey")
				fmt.Fprintf(w, "{\"status\":\"ok\",\"result\":%s}", value)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	out := fmt.Sprintf("{\"greetings\":\"hello\",\"name\":\"%s\"}", name)
	w.Header().Set("myKey", out)
}

func StartServer(t time.Duration) {
	_ = t
	http.HandleFunc("/", RPC(SetDefaultName(Sanitize(HelloHandler))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer(time.Millisecond)
}
