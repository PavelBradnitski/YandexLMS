package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	//d := time.Now().Add(100 * time.Minute)
	api, err := fetchAPI(ctx, "https://api.nbrb.by/exrates/rates?periodicity=0", 10*time.Second)
	if err != nil {
		fmt.Printf("Err: %v", err)
	} else {
		fmt.Printf("Code %v\n", api.StatusCode)
		//fmt.Printf("Body %v\n", api.Data)
	}
}

type APIResponse struct {
	Data       string // тело ответа
	StatusCode int    // код ответа
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	t := time.Now().Add(timeout)
	ctxWithDeadLine, cancel := context.WithDeadline(ctx, t)

	defer cancel()
	data := make([]byte, 4096)
	window := make([]byte, 0)
	for {
		select {
		case <-ctxWithDeadLine.Done():
			return nil, ctxWithDeadLine.Err()
		default:
			//time.Sleep(5 * time.Second)
			// fmt.Println(ctxWithDeadLine.Deadline())
			r, err := http.Get(url)
			if err != nil {
				return nil, err
			}

			select {
			case <-ctxWithDeadLine.Done():
				return nil, ctxWithDeadLine.Err()
			default:
				n, err := r.Body.Read(data)
				if err != nil && !errors.Is(err, io.EOF) {
					return nil, err
				}
				if n > 0 {
					window = append(window, data[:n]...)
				}
				if err == io.EOF {
					break
				}
			}
			return &APIResponse{Data: string(window), StatusCode: r.StatusCode}, nil
		}
	}
}
