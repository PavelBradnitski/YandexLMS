package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1*1*1+1*2"
	input1 := "(10*(10+11)*1-11)/1"
	_ = input1
	_ = input
	out, err := Calculator(input1)
	fmt.Println(out, err)
}

func Calculator(expression string) (float64, error) {
	var num float64
	var err error
	index := strings.IndexRune(expression, '(')
	if index != -1 {
		indexRight := strings.IndexRune(expression, ')')
		if indexRight == -1 {
			return 0, errors.New("bracket error")
		}
		leftIndex := findMaxIndex(expression, '(')
		num, err = Calculator(expression[findMaxIndex(expression, '(')+1 : indexRight+index])
		if err != nil {
			return 0, err
		}
		expression = expression[:leftIndex] + strconv.FormatFloat(num, 'g', -1, 64) + expression[indexRight+1:]
		return Calculator(expression)
	}

	index = strings.IndexAny(expression, "/*")
	if index != -1 {
		expression, err = simplifyExpression(expression, "*/")
		if err != nil {
			return 0, err
		}
	}

	index = strings.IndexAny(expression, "+-")
	if index != -1 {
		expression, err = simplifyExpression(expression, "+-")
		if err != nil {
			return 0, err
		}
	}

	return calculate(expression)
}

func simplifyExpression(expression, str string) (string, error) {
	index := strings.IndexAny(expression, str)
	if index != -1 {
		leftIndex := findLastIndexRune(expression[:index])
		if indexRight := strings.IndexAny(expression[index+1:], "/*+-"); indexRight != -1 {
			if leftIndex == -1 {
				num, err := calculate(expression[:indexRight+index+1])
				if err != nil {
					return "", err
				}
				expression = strconv.FormatFloat(num, 'g', -1, 64) + expression[indexRight+index+1:]
				return simplifyExpression(expression, str)
			} else {
				num, err := calculate(expression[leftIndex+1 : index+indexRight+1])
				if err != nil {
					return "", err
				}
				expression = expression[:leftIndex+1] + strconv.FormatFloat(num, 'g', -1, 64) + expression[index+indexRight+1:]
				return simplifyExpression(expression, str)
			}
		} else {
			if leftIndex != -1 {
				num, err := calculate(expression[leftIndex+1:])
				if err != nil {
					return "", err
				}
				expression = expression[:leftIndex+1] + strconv.FormatFloat(num, 'g', -1, 64)
				return simplifyExpression(expression, str)
			}
		}
	}
	return expression, nil
}

func calculate(expression string) (float64, error) {
	operators := map[rune]func(float64, float64) float64{
		'+': func(num1, num2 float64) float64 { return num1 + num2 },
		'-': func(num1, num2 float64) float64 { return num1 - num2 },
		'*': func(num1, num2 float64) float64 { return num1 * num2 },
		'/': func(num1, num2 float64) float64 { return num1 / num2 },
	}
	for operator, operation := range operators {
		index := strings.IndexRune(expression, operator)
		if index != -1 {
			num1, err := strconv.ParseFloat(expression[:index], 64)
			if err != nil {
				return 0, err
			}
			num2, err := strconv.ParseFloat(expression[index+1:], 64)
			if err != nil {
				return 0, err
			}
			if num2 == 0 && operator == '/' {
				return 0, errors.New("dividing by 0")
			}
			return operation(num1, num2), nil
		}
	}

	return 0, fmt.Errorf("неверный формат выражения")
}

func findLastIndexRune(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if strings.ContainsAny(s[i:i+1], "/*+-") {
			return i
		}
	}
	return -1
}

func findMaxIndex(str string, symbol rune) int {
	maxIndex := -1
	for i, char := range str {
		if char == symbol && i > maxIndex {
			maxIndex = i
		}
	}
	return maxIndex
}
