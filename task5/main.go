package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode/utf8"
)

type Person struct {
	name    string
	age     int
	address string
}
type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}
type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}
type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}
type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {
	return len(t.notes)
}
func (t ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, v := range t.tasks {
		if v.IsTopPriority() {
			count++
		}
	}
	return count
}

func (t ToDoList) CountOverdueTasks() int {
	count := 0
	for _, v := range t.tasks {
		if v.IsOverdue() {
			count++
		}
	}
	return count
}
func (t Task) IsOverdue() bool {
	return time.Now().Compare(t.deadline) != -1
}

func (t Task) IsTopPriority() bool {
	if t.priority == 4 || t.priority == 5 {
		return false
	}
	return true
}
func (s Student) IsExcellentStudent() bool {
	return float64(s.solvedProblems)*s.scoreForOneTask > s.passingScore
}
func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	fmt.Printf("Total Salary: %.2f\n", e.bonus+e.salary)
}
func (p Person) Print() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Address: %s\n", p.address)
}

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Animal interface {
	MakeSound()
}

type Dog struct {
}
type Cat struct {
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}
func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Logger interface {
	Log()
}
type Log struct {
	Level LogLevel
}

func (l Log) Log(str string) {
	if l.Level == Error {
		fmt.Println("ERROR:" + str)
	} else if l.Level == Info {
		fmt.Println("INFO:" + str)
	}
}
func ConcatStringsAndInt(str1, str2 string, num int) {
	fmt.Println(str1 + str2 + strconv.Itoa(num))
}

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return float64(float64(a) / float64(b)), nil
}

func GetCharacterAtPosition(str string, position int) (rune, error) {
	if utf8.RuneCountInString(str) < position {
		return 0, fmt.Errorf("position out of range")
	}
	return []rune(str)[position], nil
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	if n == 1 {
		return 1, nil
	}
	f, _ := Factorial(n - 1)
	return n * f, nil
}

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("negative numbers are not allowed")
	}
	return strconv.FormatInt(int64(num), 2), nil
}

func SumTwoIntegers(a, b string) (int, error) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	return num1 + num2, nil
}
func main() {
	// todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
	// fmt.Println(todo.TasksCount())
	// fmt.Println(todo.NotesCount())
	// fmt.Println(todo.CountTopPrioritiesTasks())
	// fmt.Print(todo.CountOverdueTasks())
	// errorLog := &Log{Level: Error}
	// errorLog.Log("This is an error message")
	//fmt.Println(DivideIntegers(4, 3))
	fmt.Println(SumTwoIntegers("a", "3"))
}
