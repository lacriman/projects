package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the ATM")

	checkPin()
	balance := readPositiveInt("Write your balance: ")

	fmt.Printf("Your balance is: %d", balance)
}

func checkPin() {
	pin := 1234
	attempts := 3

	for attempts > 0 {
		guess := readPositiveInt("Please write your PIN code: ")

		if guess == pin {
			fmt.Println("PIN correct.")
			return
		}
		attempts--
		if attempts == 0 {
			fmt.Println("You have no attempts left, the card is blocked.")
			os.Exit(1)
		} else {
			fmt.Printf("Wrong! You have %d more attempts. \n", attempts)
		}

	}

}

func readPositiveInt (message string) int {
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
