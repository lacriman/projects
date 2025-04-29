package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the divide program")

	firstNum := getNumber("Enter the first number")
	secondNum := getNumber("Enter the second number")
	for {
		if secondNum == 0 {
			fmt.Fprint(os.Stderr, "You can't divide by zero\n")
			secondNum = getNumber("Write your second number")
			continue
		}
		break
	}

	fmt.Printf("%d divide by %d is %d\n", firstNum, secondNum, firstNum/secondNum)
}

func getNumber(message string) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println(message)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)

		if input == "" {
			fmt.Fprint(os.Stderr, "Input cannot be empty\n")
			continue
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wrong input, error: %v\n", err)
			continue
		}
		return num
	}
}
