package main

import (
	"fmt"
)

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}
func (a *Account) GetBalance() float64 {
	return a.balance
}
func (a *Account) SetBalance(balance float64) error {
	if balance < 0 {
		return fmt.Errorf("can't be negative balance")
	}
	a.balance = balance
	return nil
}
func (a *Account) Deposit(balance float64) error {
	if balance < 0 {
		return fmt.Errorf("can't be negative balance")
	}
	a.balance += balance
	return nil
}
func (a *Account) Withdraw(balance float64) error {
	if balance < 0 {
		return fmt.Errorf("can't be negative balance")
	}
	if a.balance-balance > 0 {
		a.balance -= balance
	} else {
		return fmt.Errorf("can't be negative balance")
	}
	return nil
}

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}
type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string, age int) *User {
	if age == 0 {
		age = 18
	}
	return &User{Name: name, Age: age, Active: true}
}
func NewBook(title, author string, year int, genre string) *Book {
	return &Book{Title: title, Author: author, Year: year, Genre: genre}
}
func main() {

}
