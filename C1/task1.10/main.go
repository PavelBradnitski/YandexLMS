package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// logger := NewOrderLogger()
	// order := Order{1, "Иванов", 100.50}
	// logger.AddOrder(order)
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println(message)
	return nil
}

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}
type OrderLogger struct{}

func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}
func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	out := fmt.Sprintf("{\"name\":\"%s\"}", name)
	log.Print(out)
	fmt.Fprint(w, out)
}
