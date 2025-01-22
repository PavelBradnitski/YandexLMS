package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

func main() {
	// контекст, который можно отменить
	ctx := context.Background()
	ctxWithCancel, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := readSource(ctxWithCancel); err != nil {
			//при ошибке в функции чтения подадим сигнал через контекст
			cancelCtx()
			fmt.Printf("readSource(ctxWithCancel): %s\n", err)
		}
	}()
	go func() {
		defer wg.Done()
		if b, err := Contains(ctxWithCancel, bytes.NewReader([]byte("Hello, World!")), []byte("ello")); err != nil {
			fmt.Printf("processSourceData(ctxWithCancel): %s\n", err)
		} else {
			fmt.Printf("Bool: %v\n", b)
		}
	}()
	wg.Wait()
}

func readSource(ctx context.Context) error {
	// имитируем долгую работу функции
	time.Sleep(1 * time.Second)
	// допустим, возникла ошибка в процессе
	return fmt.Errorf("some error in readSource")
}

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, errors.New("sequence cannot be empty")
	}
	data := make([]byte, 4096)
	window := make([]byte, 0, len(seq))
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			n, err := r.Read(data)
			if err != nil && !errors.Is(err, io.EOF) {
				return false, err
			}
			if n > 0 {
				window = append(window, data[:n]...)
			}
			if bytes.Contains(window, seq) {
				return true, nil
			}

			if err == io.EOF {
				return false, nil
			}

			if len(window) > len(seq) {
				window = window[len(window)-len(seq):]
			}
		}
	}
}
