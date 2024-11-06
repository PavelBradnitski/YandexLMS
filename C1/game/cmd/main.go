package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"sc.io/quote/YandexLMS/C1/game/internal/application"
)

func main() {
	fmt.Println(run())
	ctx := context.Background()
	// Exit завершает программу с заданным кодом
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  10,
		Height: 10,
	}
	app := application.New(cfg)
	// Запускаем приложение
	if err := app.Run(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Println("Processing cancelled.")
		default:
			log.Println("Application run error", err)
		}
		// Возвращаем значение, не равное нулю, чтобы обозначить ошибку
		return 1
	}
	// Выход без ошибок
	return 0
}

func run() error {
	args := os.Args[1:]

	for _, v := range args {
		n, err := strconv.Atoi(v)
		if err != nil || n <= 0 {
			if err != nil {
				return err
			} else {
				return fmt.Errorf("negative number")
			}

		}
	}
	file, err := os.Create("config.txt")
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%sx%s %s%%", args[0], args[1], args[2])
	file.WriteString(str)
	defer file.Close()
	return nil
}
