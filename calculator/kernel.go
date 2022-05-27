package calculator

import (
	"GoCalc/utils"
	"strconv"
)

var numberStack utils.Stack
var signStack utils.Stack

func process(tokens *[]string) float64 {
	for i := 0; i < len(*tokens); i++ {
		token := (*tokens)[i]

		if utils.IsNumber(&token) {
			value, _ := strconv.ParseFloat(token, 0)
			numberStack.Push(value)
		} else {
			calc(token)
		}

	}

	for !signStack.IsEmpty() {
		_calc()
	}

	result, _ := numberStack.Pop()

	return result.(float64)
}

func calc(token string) {

	if signStack.IsEmpty() || token == "(" {
		signStack.Push(token)
		return
	}

	lastToken := signStack.Last().(string)

	if token == ")" {
		if lastToken != "(" {
			_calc()
			calc(token)
		} else {
			signStack.Pop()
		}
		return
	}

	if Signs[token].Priority > Signs[lastToken].Priority {
		signStack.Push(token)
	}

	if Signs[lastToken].Priority == Signs[token].Priority || Signs[token].Priority < Signs[lastToken].Priority {
		_calc()
		calc(token)
	}
}

func _calc() {
	token, _ := signStack.Pop()

	x, _ := numberStack.Pop()
	y, _ := numberStack.Pop()

	xVal := x.(float64)
	yVal := y.(float64)

	switch token {
	case "+":
		numberStack.Push(yVal + xVal)
	case "-":
		numberStack.Push(yVal - xVal)
	case "*":
		numberStack.Push(yVal * xVal)
	case "/":
		numberStack.Push(yVal / xVal)
	}
}
