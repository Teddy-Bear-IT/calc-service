package calculator

import (
	"errors"
	"strconv"
	"unicode"
)


func Calc(expression string) (float64, error) {
        tokens, err := tokenize(expression)
        if err != nil {
                return 0, err
        }
        rpn, err := toRPN(tokens)
        if err != nil {
                return 0, err
        }
        return evaluateRPN(rpn)
}

func tokenize(expression string) ([]string, error) {
        var tokens []string
        var number string

        for _, char := range expression {
                if unicode.IsDigit(char) || char == '.' {
                        number += string(char)
                } else {
                        if number != "" {
                                tokens = append(tokens, number)
                                number = ""
                        }
                        if char == ' ' {
                                continue
                        }
                        tokens = append(tokens, string(char))
                }
        }
        if number != "" {
                tokens = append(tokens, number)
        }
        return tokens, nil
}

// Преобразование выражение в обратную польскую нотацию
func toRPN(tokens []string) ([]string, error) {
        var output []string
        var stack []string
        precedence := map[string]int{
                "+": 1,
                "-": 1,
                "*": 2,
                "/": 2,
                "(": 0,
        }

        for _, token := range tokens {
                switch token {
                case "+", "-", "*", "/":
                        for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
                                output = append(output, stack[len(stack)-1])
                                stack = stack[:len(stack)-1]
                        }
                        stack = append(stack, token)
                case "(":
                        stack = append(stack, token)
                case ")":
                        for len(stack) > 0 && stack[len(stack)-1] != "(" {
                                output = append(output, stack[len(stack)-1])
                                stack = stack[:len(stack)-1]
                        }
                        if len(stack) == 0 {
                                return nil, errors.New("mismatched parentheses")
                        }
                        stack = stack[:len(stack)-1]
                default:
                        output = append(output, token)
                }
        }
        for len(stack) > 0 {
                if stack[len(stack)-1] == "(" {
                        return nil, errors.New("mismatched parentheses")
                }
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
        }
        return output, nil
}

// Вычисление обратной польской нотации :)
func evaluateRPN(rpn []string) (float64, error) {
        var stack []float64

        for _, token := range rpn {
                switch token {
                case "+", "-", "*", "/":
                        if len(stack) < 2 {
                                return 0, errors.New("invalid expression")
                        }
                        b := stack[len(stack)-1]
                        a := stack[len(stack)-2]
                        stack = stack[:len(stack)-2]
                        var result float64
                        switch token {
                        case "+":
                                result = a + b
                        case "-":
                                result = a - b
                        case "*":
                                result = a * b
                        case "/":
                                if b == 0 {
                                        return 0, errors.New("division by zero")
                                }
                                result = a / b
                        }
                        stack = append(stack, result)
                default:
                        value, err := strconv.ParseFloat(token, 64)
                        if err != nil {
                                return 0, err
                        }
                        stack = append(stack, value)
                }
        }
        if len(stack) != 1 {
                return 0, errors.New("invalid expression")
        }
        return stack[0], nil
}