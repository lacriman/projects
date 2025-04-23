package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// pin := 1234
	// attemps := 0

	fmt.Println("Welcome to the ATM")

	balance := getValidNum("Write your balance: ")

	fmt.Printf("Your balance is: %d", balance)
}

func getValidNum(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)

		if err != nil || num <= 0 {
			fmt.Println("Wrong input, please write a positive number.")
			continue			
		}

		return num
	}
}

// func withdraw() {

// }

// func deposit() {

// }
