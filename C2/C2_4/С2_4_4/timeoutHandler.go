package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	go func() {
		http.HandleFunc("/provideData", longHanlder)

		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Fatalf("error when starting a server")
		}
	}()
	StartServer(2 * time.Second)
}

func StartServer(maxTimeout time.Duration) {
	http.Handle("/readSource", http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		req, err := http.NewRequest("Get", "http://localhost:8081/provideData", nil)
		if err != nil {
			http.Error(w, "Error!", http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Request error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Reading error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}), maxTimeout, ""))
	http.ListenAndServe(":8080", nil)
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "Sleep")
}
