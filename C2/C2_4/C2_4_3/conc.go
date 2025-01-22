package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// ctx := context.Background()
	// //d := time.Now().Add(100 * time.Minute)
	// urls := []string{"https://httpbin.org/headers", "https://httpbin.org/get"}
	// apis := FetchAPI(ctx, urls, 1*time.Hour)
	// for _, api := range apis {
	// 	if api.Err != nil {
	// 		fmt.Printf("Err: %v", api.Err)
	// 	} else {
	// 		fmt.Printf("Code %v\n", api.Data)
	// 		//fmt.Printf("Body %v\n", api.Data)
	// 	}
	// }
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		http.HandleFunc("/hello", helloHandler)

		http.HandleFunc("/long", longHanlder)

		http.HandleFunc("/hi", hiHandler)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("error when starting a server")
		}
	}()
	// wg.Wait()
	//time.Sleep(100)
	testCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// urls := []string{"https://httpbin.org/headers", "https://httpbin.org/get"}
	urls := []string{"http://localhost:8080/hello", "http://localhost:8080/hi"}
	apis := FetchAPI(testCtx, urls, 1*time.Microsecond)
	for _, api := range apis {
		if api.Err != nil {
			fmt.Printf("Err: %v\n", api.Err)
		} else {
			fmt.Printf("Code %v\n", api.Data)
			//fmt.Printf("Body %v\n", api.Data)
		}
	}
}

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var out []*APIResponse
	mu := &sync.Mutex{}
	client := &http.Client{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			t := time.Now().Add(timeout)
			ctxWithDeadLine, cancel := context.WithDeadline(ctx, t)
			defer cancel()

			r, err := http.NewRequestWithContext(ctxWithDeadLine, http.MethodGet, url, nil)
			if err != nil {
				mu.Lock()
				out = append(out, &APIResponse{Err: err})
				mu.Unlock()
				return
			}

			select {
			case <-ctxWithDeadLine.Done():
				mu.Lock()
				out = append(out, &APIResponse{Err: ctxWithDeadLine.Err()})
				mu.Unlock()
				return
			default:

				resp, err := client.Do(r)
				if err != nil {
					mu.Lock()
					out = append(out, &APIResponse{Err: err})
					mu.Unlock()
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil && !errors.Is(err, io.EOF) {
					mu.Lock()
					out = append(out, &APIResponse{Err: err})
					mu.Unlock()
					return
				}
				mu.Lock()
				out = append(out, &APIResponse{URL: url, Data: string(body), StatusCode: resp.StatusCode, Err: err})
				mu.Unlock()
			}
		}(url)
	}
	wg.Wait()
	return out
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Microsecond * 100)
	fmt.Fprintf(w, "Hi, World!")
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}
