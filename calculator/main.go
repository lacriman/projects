package main

import "fmt"

func main() {
	num1 := getValidInput("Welcome to a calculator, write your first number: ")
	num2 := getValidInput("Write your second number: ")
	operator := getValidOperator("Write an operator: +, -, * or /")

	switch operator {
	case "+":
		res := num1 + num2
		fmt.Println(num1, " + ", num2, " = ", res)

	case "-":
		res := num1 - num2
		fmt.Println(num1, " - ", num2, " = ", res)

	case "*":
		res := num1 * num2
		fmt.Println(num1, " * ", num2, " = ", res)

	case "/":
		res := num1 / num2
		fmt.Println(num1, " / ", num2, " = ", res)
	}
}

func getValidInput(message string) (num int) {
	fmt.Println(message)
	fmt.Scan(&num)
	return
}

func getValidOperator(message string) (operator string) {
	fmt.Println(message)
	fmt.Scan(&operator)

	for {
		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			break
		}
		fmt.Println("Wrong input, write an operator: +, -, * or /")
		fmt.Scan(&operator)
	}
	return
}
