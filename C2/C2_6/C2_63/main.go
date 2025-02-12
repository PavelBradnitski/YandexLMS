package main

import (
	"fmt"
	"time"
)

func main() {
	// // questions := []string{"test1", "test2", "test3"}
	// // answers := []string{"1", "2", "3"}
	// // answerCh := make(chan string, 1)
	// // fmt.Println(QuizRunner(questions, answers, answerCh))
	fillAnswers := func(answers []string, answerCh chan string) {
		go func() {
			for i := range answers {
				answerCh <- answers[i]
			}
		}()
	}

	checkAnswers := func(questions []string, answers []string, expectedCorrectAnswers int, answerCh chan string) {
		correctAnswers := QuizRunner(questions, answers, answerCh)
		fmt.Println(correctAnswers)
		if correctAnswers != expectedCorrectAnswers {
			fmt.Printf("Expected %d correct answers, but got %d", expectedCorrectAnswers, correctAnswers)
		}
	}

	questions := []string{
		"What is 2 + 2 ?",
		"What is 2 * 2 ?",
		"What is 2 - 2 ?",
	}
	faileAnswers := []string{"4", "3", "0"}
	rightAnswers := []string{"4", "4", "0"}

	answerCh := make(chan string)
	defer close(answerCh)
	go fillAnswers(faileAnswers, answerCh)
	checkAnswers(questions, rightAnswers, 2, answerCh)

	go fillAnswers(rightAnswers, answerCh)
	checkAnswers(questions, rightAnswers, 3, answerCh)

	checkAnswers(questions, rightAnswers, 0, answerCh)
}
func QuizRunner(questions, answers []string, answerCh <-chan string) int {
	out := 0
	for i := range questions {
		select {
		case answer := <-answerCh:
			if answer == answers[i] {
				out++
			}
		case <-time.After(1 * time.Second):
			continue
		}
	}
	return out
}
